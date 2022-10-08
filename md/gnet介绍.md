#开篇

上一篇Go netpoletpl大解析我们分析了Go原生网络模型以及部分源码，绝大部分场景下(99%)，使用原生netpoll已经足够了。

但是在一些海量并发连接下，原生netpoll会为每一个连接都开启一个goroutine处理，也就是1千万的连接就会创建一千万个goroutine。
这就给了这些特殊场景下的优化空间，这也是像gnet和cloudwego/netpoll诞生的原因之一吧。

本质上他们的底层核心都是一样的，都是基于epoll(linux)实现的。只是事件发生后，每个库的处理方式会有所不同。

本篇文章主要分析gnet的。至于使用姿势就不发了，gnet有对应的demo库，可以自行体验。

#架构
直接引用gnet官网的一张图

![图片](./image/gnet官方图片.png)

gnet采用的是『主从多 Reactors』。也就是一个主线程负责监听端口连接，当一个客户端连接到来时，就把这个连接根据负载均衡算法分配给其中一个sub线程，由对应的sub线程去处理这个连接的读写事件以及管理它的死亡。

下面这张图就更清晰了。
![图片](./image/gnet执行流程.png)
## 核心结构
我们先解释gnet的一些核心结构。

```go
type engine struct {
	ln           *listener          // the listener for accepting new connections
	lb           loadBalancer       // event-loops for handling events
	wg           sync.WaitGroup     // event-loop close WaitGroup
	opts         *Options           // options with engine
	once         sync.Once          // make sure only signalShutdown once
	cond         *sync.Cond         // shutdown signaler
	mainLoop     *eventloop         // main event-loop for accepting connections
	inShutdown   int32              // whether the engine is in shutdown
	tickerCtx    context.Context    // context for ticker
	cancelTicker context.CancelFunc // function to stop the ticker
	eventHandler EventHandler       // user eventHandler
}
```
engine就是程序最上层的结构了。

    ln对应的listener就是服务启动后对应监听端口的监听器。
    
    lb对应的loadBalancer就是负载均衡器。也就是当客户端连接服务时，负载均衡器会选择一个sub线程，把连接交给此线程处理。
    
    mainLoop 就是我们的主线程了，对应的结构eventloop。当然我们的sub线程结构也是eventloop。结构相同，不同的是职责。
        主线程负责的是监听端口发生的客户端连接事件，然后再由负载均衡器把连接分配给一个sub线程。而sub线程负责的是绑定分配给他的连接(不止一个)，
        且等待自己管理的所有连接后续读写事件，并进行处理。
    
接着看eventloop。

![图片](./image/eventloop.png)

    netpoll.Poller:每一个 eventloop都对应一个epoll或者kqueue。
    
    buffer用来作为读消息的缓冲区。
    
    connCoun记录当前eventloop存储的tcp连接数。
    
    udpSockets和connetcions分别管理着这个eventloop下所有的udp socket和tcp连接，注意他们的结构map。这里的int类型存储的就是fd。



对应conn结构，

```go
type conn struct {
	fd             int                     // file descriptor
	ctx            interface{}             // user-defined context
	peer           unix.Sockaddr           // remote socket address
	loop           *eventloop              // connected event-loop
	buffer         []byte                  // buffer for the latest bytes
	opened         bool                    // connection opened event fired
	localAddr      net.Addr                // local addr
	remoteAddr     net.Addr                // remote addr
	isDatagram     bool                    // UDP protocol
	inboundBuffer  elastic.RingBuffer      // buffer for leftover data from the peer
	outboundBuffer *elastic.Buffer         // buffer for data that is eligible to be sent to the peer
	pollAttachment *netpoll.PollAttachment // connection attachment for poller
}
```

这里面有几个字段介绍下，

    buffer:存储当前conn对端(client)发送的最新数据，比如发送了三次，那个此时buffer存储的是第三次的数据,代码里有。
    inboundBuffer:存储对端发送的且未被用户读取的剩余数据，还是个Ring Buffer。
    outboundBuffer:存储还未发送给对端的数据。(比如服务端响应客户端的数据，由于conn fd是不阻塞的，调用write返回不可写的时候，就可以先把数据放到这里)

conn相当于每个连接都会有自己独立的缓存空间。这样做是为了减少集中式管理内存带来的锁问题。使用Ring buffer是为了增加空间的复用性。

整体结构就这些。

## 核心逻辑
当程序启动时，

```go
func serve(eventHandler EventHandler, listener *listener, options *Options, protoAddr string) error {
    // Figure out the proper number of event-loops/goroutines to run.
    numEventLoop := 1
    if options.Multicore {
    numEventLoop = runtime.NumCPU()
    }
    if options.NumEventLoop > 0 {
    numEventLoop = options.NumEventLoop
    }
	// ...
}
```

    会根据用户设置的options明确eventloop循环的数量，也就是有多少个sub线程。再进一步说，在linux环境就是会创建多少个epoll对象。
    那么整个程序的epoll对象数量就是count(sub)+1(main Listener)。

```go
func (eng *engine) activateReactors(numEventLoop int) error {
	for i := 0; i < numEventLoop; i++ {
		if p, err := netpoll.OpenPoller(); err == nil {
			el := new(eventloop)
			el.ln = eng.ln
			el.engine = eng
			el.poller = p
			el.buffer = make([]byte, eng.opts.ReadBufferCap)
			el.connections = make(map[int]*conn)
			el.eventHandler = eng.eventHandler
			eng.lb.register(el)
		} else {
			return err
		}
	}

	// Start sub reactors in background.
	eng.startSubReactors()

	if p, err := netpoll.OpenPoller(); err == nil {
		el := new(eventloop)
		el.ln = eng.ln
		el.idx = -1
		el.engine = eng
		el.poller = p
		el.eventHandler = eng.eventHandler
		if err = el.poller.AddRead(eng.ln.packPollAttachment(eng.accept)); err != nil {
			return err
		}
		eng.mainLoop = el

		// Start main reactor in background.
		eng.wg.Add(1)
		go func() {
			el.activateMainReactor(eng.opts.LockOSThread)
			eng.wg.Done()
		}()
	} else {
		return err
	}

	// Start the ticker.
	if eng.opts.Ticker {
		go eng.mainLoop.ticker(eng.tickerCtx)
	}

	return nil
}
```

上图就是我说的，会根据设置的数量创建对应的eventloop,把对应的eventloop 注册到负载均衡器中。
当新连接到来时，就可以根据一定的算法(gnet提供了轮询、最少连接以及hash)挑选其中一个eventloop把连接分配给它。
我们先来看主线程，(由于我使用的是mac,所以后面关于IO多路复用，实现部分就是kqueue代码了，当然原理是一样的)


Polling就是等待网络事件到来，传递了一个闭包参数，更确切的说是一个事件到来时的回调函数，从名字可以看出，就是处理新连接的。



至于Polling函数，

图片

逻辑很简单，一个for循环等待事件到来，然后处理事件。

主线程的事件分两种，

一种是正常的fd发生网络连接事件，

一种是通过NOTE_TRIGGER立即激活的事件。

图片

通过NOTE_TRIGGER触发告诉你队列里有task任务，去执行task任务。

如果是正常的网络事件到来，就处理闭包函数，主线程处理的就是上面的accept连接函数。


accept连接逻辑很简单，拿到连接的fd。设置fd非阻塞模式(想想连接是阻塞的会咋么样?),然后根据负载均衡算法选择一个sub 线程，通过register函数把此连接分配给它。


register做了两件事，首先需要把当前连接注册到当前sub 线程的epoll or kqueue 对象中,新增read的flag。



接着就是把当前连接放入到connections的map结构中 fd->conn。



这样当对应的sub线程事件到来时，可以通过事件的fd找到是哪个连接，进行相应的处理。


如果是可读事件，


到这里分析差不多就结束了。



总结
在gnet里面，你可以看到，基本上所有的操作都无锁的。



那是因为事件到来时，采取的都是非阻塞的操作，且是串行处理对应的每个fd(conn)。每个conn操作的都是自身持有的缓存空间。同时处理完一轮触发的所有事件才会循环进入下一次等待，在此层面上解决了并发问题。



当然这样用户在使用的时候也需要注意一些问题，比如用户在自定义EventHandler中，如果要异步处理逻辑，就不能像下面这样开一个g然后在里面获取本次数据，


而应该先拿到数据，再异步处理。


issues上有提到，连接是使用map[int]*conn存储的。gnet本身的场景就是海量并发连接，内存会很大。进而big map存指针会对 GC造成很大的负担，毕竟它不像数组一样，是连续内存空间，易于GC扫描。



还有一点，在处理buffer数据的时候，就像上面看到的，本质上是将buffer数据copy给用户一份，那么就存在大量copy开销,在这一点上，字节的netpoll实现了Nocopy Buffer，改天研究一下。


推荐阅读

Go netpoll大解析



福利

我为大家整理了一份从入门到进阶的Go学习资料礼包，包含学习建议：入门看什么，进阶看什么。关注公众号 「polarisxu」，回复 ebook 获取；还可以回复「进群」，和数万 Gopher 交流学习。

阅读 3033
收藏此内容的人还喜欢
国外大佬的 4 个项目 yyds
逛逛GitHub
赞 31
不喜欢
不看的原因
内容质量低不看此公众号

一款强大的在线编译器工具，支持 38 种开发语言
小集
赞 11
不喜欢
不看的原因
内容质量低不看此公众号

百亿级数据 分库分表 后怎么分页查询？
码猿技术专栏
阅读 3032
不喜欢
不看的原因
内容质量低不看此公众号


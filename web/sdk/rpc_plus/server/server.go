package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloService struct {
}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello response " + request
	return nil
}

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, srv)
}

func main() {
	RegisterHelloService(new(HelloService))

	listen, err := net.Listen("tcp", ":2234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		//go rpc.ServeConn(accept)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(accept)) //使用json返回数据
	}

}

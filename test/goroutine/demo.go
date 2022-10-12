package main

import (
	"fmt"
	"sync"
	"time"
)

// go build -race demo.go 可以查看是否存在资源竞争问题

// 计算1-200的阶乘 使用goroutine实现

var (
	result = make(map[int]int64, 16)
	lock   sync.Mutex
)

func Test(n int) {

	var res int64 = 0
	for i := 1; i <= n; i++ {
		res += int64(i)
	}

	// 结果存到map中
	lock.Lock()
	result[n] = res //concurrent map writes 存在资源竞争问题
	lock.Unlock()
}

func main() {
	// 开启多个协程
	for i := 1; i <= 200; i++ {
		go Test(i)
	}

	time.Sleep(time.Second * 10)
	// 遍历map
	for k, v := range result {
		fmt.Printf("map[%d]=%d\n", k, v)
	}
}

package main

import (
	"fmt"
	"time"
)

// 一个协程往管道中写数据 一个协程从管道中读数据
func WriteData(intChain chan int) {
	defer close(intChain)
	for i := 1; i < 100; i++ {
		intChain <- i
		fmt.Println("往管道中写入数据：", i)
		time.Sleep(time.Nanosecond * 1) // 休眠50毫秒
	}
}

func ReadData(intChain chan int, existChain chan bool) {
	for {
		v, ok := <-intChain
		if !ok {
			break
		}
		fmt.Println("从管道中读取数据：", v)
	}
	existChain <- true
	close(existChain)
}

func main() {
	intChain := make(chan int, 10)
	existChain := make(chan bool, 1)
	go WriteData(intChain)
	go ReadData(intChain, existChain)
	for {
		_, ok := <-existChain
		if ok {
			break
		}
	}

}

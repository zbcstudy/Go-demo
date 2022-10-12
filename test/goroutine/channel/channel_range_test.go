package main

import (
	"fmt"
	"testing"
)

// channel遍历
func TestChannelRange(t *testing.T) {
	intChain := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChain <- i * i
	}
	// 管道关闭 必须
	close(intChain)
	// 遍历时如果管道没有关闭 则会出现deadlock错误
	// 遍历时如果管道已经关闭 则可以正常遍历数据 遍历结束之后 退出遍历
	for v := range intChain {
		fmt.Println(v)
	}
}

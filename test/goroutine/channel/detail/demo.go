package main

import (
	"fmt"
)

func main() {
	intChain := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		intChain <- i
	}

	stringChain := make(chan string, 10)
	for i := 1; i <= 10; i++ {
		stringChain <- "hello" + fmt.Sprintf("%v", i)
	}
	// 传统的方法在管道遍历时 如果不关闭会阻塞 导致deadlock
	for {
		select {
		case v := <-intChain:
			fmt.Printf("从int chan中取出数据：%d\n", v)
		case v := <-stringChain:
			fmt.Printf("从string chan中取出数据：%v\n", v)
		default:
			fmt.Printf("都取不到数据\n")
			return
		}
	}
}

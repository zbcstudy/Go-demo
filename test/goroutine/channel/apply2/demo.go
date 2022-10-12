package main

import (
	"fmt"
	"time"
)

// 求取素数
func main() {
	intChain := make(chan int, 1000)
	primeChain := make(chan int, 1000) //结果chain

	existChain := make(chan bool, 4)

	go PutNum(&intChain)

	for i := 0; i < 4; i++ {
		go PrimeNum(&intChain, primeChain, existChain)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-existChain
		}
		// 当我们从existChain中取出4个数据 就可以放心的关闭primeChain
		close(primeChain)
	}()

	// 遍历取出数据
	for {
		re, ok := <-primeChain
		if !ok {
			break
		}
		fmt.Println("素数：", re)
	}
	fmt.Println("主线程退出")
}

func PutNum(intChain *chan int) {
	for i := 1; i <= 1000; i++ {
		*intChain <- i
	}
	close(*intChain)
	time.Sleep(time.Second * 3)
}

func PrimeNum(intChain *chan int, primeChain chan int, existChain chan bool) {
	var flag bool
	for {
		num, ok := <-*intChain
		if !ok { //未取到
			break
		}
		flag = true
		for i := 2; i <= num-1; i++ {
			if num%i == 0 {
				// 不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChain <- num
		}
	}

	fmt.Println("有一个协程取不到数据退出了")

	existChain <- true
}

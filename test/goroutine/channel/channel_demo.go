package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Println(intChan)

	// 向管道写入数据
	intChan <- 10

	num := 211
	intChan <- num

	intChan <- 20

	fmt.Println("len:", len(intChan)) // 长度
	fmt.Println("cap:", cap(intChan)) // 容量

	// 从管道中拿数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2:", num2)
	fmt.Println("len:", len(intChan)) // 长度
	fmt.Println("cap:", cap(intChan)) // 容量
}

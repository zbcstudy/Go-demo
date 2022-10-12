package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCPU := runtime.NumCPU()
	fmt.Println("numCPU:", numCPU)

	// 自己设置使用多少CPU
	runtime.GOMAXPROCS(numCPU - 1)

	numGoroutine := runtime.NumGoroutine()
	fmt.Println("numGoroutine:", numGoroutine)

}

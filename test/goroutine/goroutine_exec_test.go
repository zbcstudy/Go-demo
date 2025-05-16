package main

import (
	"fmt"
	"testing"
	"time"
)

func TestExec(t *testing.T) {

	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine run i=%d\n", i)
			time.Sleep(1 * time.Second)
		}
	}()

	i := 1
	for {
		i++
		fmt.Printf("main goroutine run i=%d\n", i)
		time.Sleep(1 * time.Second)
		if i == 5 {
			break
		}
	}
}

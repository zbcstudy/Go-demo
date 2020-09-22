package test

import (
	"fmt"
	"testing"
)

func TestFirst(t *testing.T) {
	t.Log("first log test")
}

func TestFibList(t *testing.T) {
	var a = 1
	var b = 1
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		fmt.Println(" ", b)
		temp := a
		a = b
		b = temp + a
	}

}

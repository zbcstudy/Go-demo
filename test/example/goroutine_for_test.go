package example

import (
	"fmt"
	"testing"
)

// 在for循环中使用goroutine
// 闭包是一个特殊的匿名函数，可以访问函数体外部的变量
// 输出的值并不是我们希望的值
func TestForWithErr(t *testing.T) {
	done := make(chan bool)

	values := []string{"a", "b", "c"}

	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	for _ = range values {
		<-done
	}
}

// 想要避免上面的情况发生 可以将变量v的值作为参数传递给闭包
func TestForWithOk1(t *testing.T) {
	done := make(chan bool)

	values := []string{"a", "b", "c"}

	for _, v := range values {
		go func(param string) {
			fmt.Printf("val=%s,pointer=%p\n", param, &param)
			done <- true
		}(v)
	}

	for _ = range values {
		<-done
	}
}

func TestForWithOk2(t *testing.T) {
	done := make(chan bool)

	values := []string{"a", "b", "c"}

	for _, v := range values {
		param := v
		go func() {
			fmt.Printf("val=%s,pointer=%p\n", param, &param)
			done <- true
		}()
	}

	for _ = range values {
		<-done
	}
}

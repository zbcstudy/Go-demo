package common

import (
	"fmt"
	"testing"
)

func TestErrorCatch(t *testing.T) {
	test()
	fmt.Println("test 执行完成")
}

func test() {
	//defer+recover 用来捕获和处理异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获到异常", err)
		}
	}()

	num1 := 10
	num2 := 10
	res := num1 / num2
	fmt.Printf("res=%v", res)

}

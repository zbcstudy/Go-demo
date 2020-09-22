package test

import (
	"fmt"
	"testing"
	time "time"
)

/**
 * 接口为非入侵性 实现不依赖于接口定义
 * 所有接口的定义可以包含在接口使用者包内
 */
type Programmer interface {
	write() string
}

type GoProgrammer struct {
}

func (gp *GoProgrammer) write() string {
	return "hello world"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.write())
}

type IntConv func(op int) int

func timeSpent(conv IntConv) IntConv {
	return func(op int) int {
		start := time.Now()
		res := conv(op)
		fmt.Println("time spent:", time.Since(start).Microseconds(), "res:", res)
		return res
	}
}

func TestTimeSpent(t *testing.T) {
	timeSpent(func(op int) int {
		return op * op
	})
}

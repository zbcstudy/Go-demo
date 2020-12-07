package _func

import (
	"fmt"
	"testing"
)

func sum(ops ...int) int {
	var res = 0
	for _, op := range ops {
		res += op
	}
	return res
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3, 4))
}

func Clear() {
	fmt.Println("last clear resource")
}

// defer 关键字用来最后资源的释放
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("start")
	panic("error")
}

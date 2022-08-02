package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandSeed(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	println(rand.Intn(100)) // 不设置随机种子 每次获取的值是一样的
}

func TestRandom(t *testing.T) {
	count := 0
	for {
		count++
		i := rand.Intn(100)
		println(i)
		if i == 99 {
			break
		}
	}
	fmt.Printf("一共执行了%v次", count)
}

// break 跳到指定标签
func TestBreakJump(t *testing.T) {
	//label2:
	for i := 0; i < 4; i++ {
	label1:
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label1
			}
			fmt.Println("j=", j)
		}
	}
}

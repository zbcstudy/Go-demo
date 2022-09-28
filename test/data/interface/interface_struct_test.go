package _interface

import (
	"fmt"
	"testing"
)

type Monkey struct {
	Name string
}

type LittleMonkey struct {
	Monkey
}

type BirdAble interface {
	Fly()
}

type FishAble interface {
	Swimming()
}

func (m *Monkey) Climbing() {
	fmt.Println(m.Name, "生来会爬树")
}

func (lm LittleMonkey) Fly() {
	fmt.Println(lm.Name, "学会了飞翔")
}

func (lm LittleMonkey) Swimming() {
	fmt.Println(lm.Name, "学会了游泳")
}

// 接口可以看成是对继承的补充
// 实现接口中方法 可以在不破坏继承的基础上 对功能进行扩展
func TestDiff(t *testing.T) {
	monkey := Monkey{Name: "lm-name"}
	lm := LittleMonkey{monkey}
	lm.Climbing()
	lm.Fly()
	lm.Swimming()
}

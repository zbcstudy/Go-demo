package main

import (
	"fmt"
	"testing"
)

type Point struct {
	x int
	y int
}

type Rect struct {
	left, right Point
}

type RectP struct {
	left, right *Point
}

func TestStructPoint(t *testing.T) {
	rect := Rect{Point{1, 2}, Point{3, 4}}
	// 地址之间相隔8个字节
	fmt.Printf("rect.left.x的地址为：%p,rect.left.y的地址为：%p,rect.right.x的地址为：%p,rect.right.y的地址为：%p\n",
		&rect.left.x,
		&rect.left.y,
		&rect.right.x,
		&rect.right.y)
	fmt.Printf("rect.left本身指向的地址：%p,rect.right本身指向的地址：%p\n", &rect.left, &rect.right)

	rectP := RectP{&Point{10, 20}, &Point{30, 40}}
	fmt.Printf("rectp.left.x的地址为：%p,rectp.left.y的地址为：%p,rectp.right.x的地址为：%p,rectp.right.y的地址为：%p\n",
		&rectP.left.x,
		&rectP.left.y,
		&rectP.right.x,
		&rectP.right.y)
	fmt.Printf("rectP.left本身指向的地址：%p,rectP.right本身指向的地址：%p", &rectP.left, &rectP.right)

}

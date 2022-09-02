package data

import (
	"fmt"
	"testing"
)

func TestArrayList(t *testing.T) {
	var arr [3]int
	t.Log(arr)
	arr1 := [4]int{1, 2, 3, 4}
	t.Log(arr1)
	//不定长数组
	arr2 := [...]int{2, 3, 4, 5}
	t.Log(arr2)
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{2, 3, 4, 5, 6}
	for index, value := range arr {
		t.Log(index, value)
	}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	t.Log("数组截取：", arr[1:3])
}

//切片
func TestSlice(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 34, 4}
	t.Log(len(s1), cap(s1))

	s3 := make([]int, 3, 5)
	t.Log(len(s3), cap(s3))
	// 只能访问len 长度的元素
	t.Log(s3[0], s3[1], s3[2])

	s3 = append(s3, 33)
	t.Log(s3[3])
}

func TestSliceGrowing(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	// a== b 不能直接使用
	t.Log(a, b)
}

func test01(arr []int) {
	fmt.Println("test01 array=", arr)
	if len(arr) > 0 {
		arr[0] = 10
	}
	fmt.Println("test01 array result=", arr)
}

func TestArrayInvoke(t *testing.T) {
	arr := []int{1, 3, 5}
	test01(arr)
	fmt.Println("TestArrayInvoke arr=", arr)
}

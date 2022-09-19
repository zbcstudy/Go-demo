package slice

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSliceDemo(t *testing.T) {
	arr := [5]int{11, 22, 33, 44, 55}
	// 截取数组 创建一个切片 起始index包含，终止index不包含
	slice := arr[1:3]
	fmt.Printf("slice的类型为：%T,值为%v\n", slice, slice)
	println("slice的长度为：", len(slice), "slice的容量为：", cap(slice))

}

func TestSliceCreate(t *testing.T) {
	// 1 定义一个切片 让切片取引用一个创建好的数组

	// 2 make()
	var slice = make([]int, 2, 4)
	println(slice[1])

	ints := append(slice, 3)
	fmt.Printf("%v\n", ints)

	// 3
	var slice2 = []int{12, 23, 45}
	fmt.Printf("slice2的类型：%T,值为：%v", slice2, slice2)
}

func TestSliceAppend(t *testing.T) {
	arr := [5]int{11, 22, 33, 44, 55}

	slice1 := arr[1:]
	fmt.Println("slice1:", slice1)

	var slice2 = []int{12, 23, 45}

	// slice追加slice
	slice1 = append(slice1, slice2...)
	fmt.Println("slice1 append is:", slice1)
}

func TestSliceCopy(t *testing.T) {
	var slice1 []int = []int{1, 23, 4, 56, 7}
	slice2 := make([]int, 10)

	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	// 调用copy时, 数据都要是切片
	copy(slice2, slice1)
	// slice1和slice2的数据空间相互独立 互不影响
	slice1[0] = 100
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)

	println("*****************************")

	// slice3长度为1
	slice3 := make([]int, 1)
	// copy时多出来的数据直接丢弃
	copy(slice3, slice1)
	fmt.Println("slice3:", slice3)
}

func TestSliceWithString(t *testing.T) {
	str := "TestSliceWithString"
	fmt.Println("str point:", &str)
	// 使用切片进行截取
	slice := str[4:]
	str = str + strconv.Itoa(1)
	fmt.Println("new str point:", &str, "value:", str)
	fmt.Printf("type:%T,slice:%v\n", slice, slice)

	// 修改字符串的首字母
	slice2 := []byte(str)
	slice2[0] = 't'
	str = string(slice2)
	fmt.Println("str:", str)
}

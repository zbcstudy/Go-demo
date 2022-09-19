package pkg

import (
	"fmt"
	"testing"
)

func BubbleSort(arr *[5]int) {
	fmt.Println("排序前：", *arr)
	temp := 0
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1; j++ {
			if arr[j] > arr[j+1] {
				// 数据交换
				temp = (*arr)[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	arr := [5]int{20, 5, 80, 30, 10}
	BubbleSort(&arr)
	fmt.Println("排序后：", arr)
}

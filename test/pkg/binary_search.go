package pkg

import "fmt"

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		// 未找到
		println("未找到")
		return
	}
	// 找到中间值
	middleIndex := (rightIndex + leftIndex) / 2

	if arr[middleIndex] > findVal {
		// 说明我们要找的在left 和 middle之间
		BinaryFind(arr, leftIndex, middleIndex-1, findVal)
	} else if arr[middleIndex] < findVal {
		BinaryFind(arr, middleIndex+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了 下标为：%v", middleIndex)
	}
}

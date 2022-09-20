package pkg

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := [6]int{1, 88, 200, 1000, 1024, 2048}
	BinaryFind(&arr, 0, len(arr), 888)
}

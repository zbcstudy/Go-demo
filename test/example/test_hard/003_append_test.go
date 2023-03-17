package test_hard

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	s2 := make([]int, 0)
	s2 = append(s2, 1, 2, 3)
	fmt.Println(s2)
}

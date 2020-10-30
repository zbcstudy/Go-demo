package data

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 1
	for n <= 5 {
		t.Log(n)
		n++
	}
}

// if 判断语句支持两端判断
func TestMultiSec(t *testing.T) {
	if a, err := someFunc(); err != nil {
		t.Log(a, true)
	} else {
		t.Log(false)
	}
}

func someFunc() (interface{}, interface{}) {
	return 1, 1
}

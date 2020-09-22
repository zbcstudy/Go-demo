package test

import "testing"

func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	set := map[int]bool{}
	set[1] = true
	n := 1
	if set[n] {
		t.Logf("%d is existed", n)
	} else {
		t.Logf("%d not existed", n)
	}
	delete(set, 1)
}

package common

import (
	"testing"
)

func TestDeferClose(t *testing.T) {
	var whatever [5]struct{}
	for i := range whatever {
		defer func() { println(i) }()
	}
}

func TestDeferClose2(t *testing.T) {
	var whatever [5]struct{}
	for i := range whatever {
		defer println(i)
	}
}

type T struct {
	string
}

func Close(i T) {
	println(i.string, " closing")
}

func TestDeferClose3(t *testing.T) {
	whatever := [3]T{{"a"}, {"b"}, {"c"}}
	for _, v := range whatever {
		defer Close(v)
	}
}

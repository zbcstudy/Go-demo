package example

import (
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	for i := 10; i < 31; i++ {
		println(i, int(math.Pow(2, float64(i))))
	}
}

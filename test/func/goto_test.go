package _func

import (
	"fmt"
	"testing"
)

func TestGoto(t *testing.T) {
	fmt.Println("ok1")
	goto label1
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
label1:
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
}

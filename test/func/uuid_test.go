package _func

import (
	"fmt"
	"github.com/hashicorp/go-uuid"
	"strings"
	"testing"
)

// github.com/hashicorp/go-uuid uuid@v3.2.0+incompatible

func TestUUID(t *testing.T) {
	str, _ := uuid.GenerateUUID()
	fmt.Println(str)
	fmt.Println(strings.ReplaceAll(str, "-", ""))
}

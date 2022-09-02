package pkg

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/rogpeppe/fastuuid"
	"strings"
	"testing"
)

func TestUUId(t *testing.T) {
	generator, _ := fastuuid.NewGenerator()
	fmt.Println(generator.Next())
	hex128 := generator.Hex128()
	rep := strings.ReplaceAll(hex128, "-", "")
	fmt.Println(hex128, rep)
}

func TestUUIDWithGoogle(t *testing.T) {
	u := uuid.New()
	println(u.ID())
	println(u.String())
	println(u.URN())
	for i := 1; i < 10; i++ {
		println(uuid.NewString())
	}
}

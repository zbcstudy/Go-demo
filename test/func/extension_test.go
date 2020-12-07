package _func

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("....")
}

func (p *Pet) SpeakTo(name string) {
	p.Speak()
	fmt.Println(" ", name)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	d.p.Speak()
}

func (d *Dog) SpeakTo(name string) {
	d.p.SpeakTo(name)
}

func TestExtension(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("wang")
}

package test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type singleton struct{}

var singletonInstance *singleton
var once sync.Once

func GetSingletonObj() *singleton {
	once.Do(func() {
		fmt.Println("create obj")
		singletonInstance = new(singleton)
	})
	return singletonInstance
}

func TestSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%d\n", unsafe.Pointer(obj))
		}()
	}
	wg.Done()
}

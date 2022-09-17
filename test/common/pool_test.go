package common

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// pool中存储的对象会在gc时进行清除
func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new object")
			return 100
		}}
	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(23)
	runtime.GC() //会人为触发一次gc
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}

// pool在多携程下的使用
func TestSyncPoolInMultiGoRoutine(t *testing.T) {
	pool := &sync.Pool{New: func() interface{} {
		fmt.Println("create new object")
		return 10
	}}
	pool.Put(100)
	pool.Put(100)
	pool.Put(100)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}

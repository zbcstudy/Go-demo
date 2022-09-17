package common

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Microsecond * 50)
}

//
func TestCounterThreadSafe(t *testing.T) {
	var mux sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func(i int) {
			defer func() {
				mux.Unlock()
			}()
			mux.Lock()
			counter++
		}(i)
	}
	time.Sleep(time.Microsecond * 50)
	t.Logf("counter= %d", counter)
}

func TestCounterThreadSafeWithGroup(t *testing.T) {
	var mux sync.Mutex
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				mux.Unlock()
			}()
			mux.Lock()
			counter++
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Logf("counter= %d", counter)
}

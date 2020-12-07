package _func

import (
	"fmt"
	"testing"
	"time"
)

func cancel1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func cancel2(cancelChan chan struct{}) {
	close(cancelChan)
}

func isCanceled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, chanCh chan struct{}) {
			for {
				if isCanceled(cancelChan) {
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i, " Done")
		}(i, cancelChan)
	}
	//cancel1(cancelChan)
	cancel2(cancelChan)
	time.Sleep(time.Second * 1)
}

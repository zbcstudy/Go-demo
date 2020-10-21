package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func hasCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancelWithContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int) {
			for {
				if hasCanceled(ctx) {
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i, " Done")
		}(i)
	}
	//cancel1(cancelChan)
	cancel()
	time.Sleep(time.Second * 1)
}

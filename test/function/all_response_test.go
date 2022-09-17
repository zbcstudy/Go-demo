package function

import (
	"runtime"
	"testing"
	"time"
)

// 所有任务完成
func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	//ch := make(chan string)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"

	}
	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("before: ", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("after: ", runtime.NumGoroutine())
}

package data

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	fmt.Printf("now=%v;类型为：%T\n", now, now)
	fmt.Printf("now=%v;类型为：%#v\n", now, now)
	fmt.Println(now.String())
	fmt.Println(now.Month())
	// 时间必须是 2006-01-02 15:04:05
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Location())
	fmt.Println(now.Local())

	fmt.Println(now.UnixNano())
	fmt.Println(now.Unix())
}

package data

import (
	"fmt"
	"strconv"
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

func TestSleep(t *testing.T) {
	i := 1
	for {
		println(i)
		time.Sleep(time.Second)
		if i > 15 {
			break
		}
		i++
	}
}

func TestUnix(t *testing.T) {
	now := time.Now()
	println("unix:", now.Unix(), "\nunixNa", now.UnixNano())
}

func TestCostTime(t *testing.T) {
	start := time.Now().Unix()
	str := ""
	for i := 0; i < 50000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
	end := time.Now().Unix()
	fmt.Printf("执行循环耗费时间：%v", end-start)

}

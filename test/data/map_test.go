package data

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	t.Log(m)
	m2 := map[string]int{}
	t.Log(m2["2"])
}

func TestTravelMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		t.Log(k, "----", v)
	}
}

func TestMapDelete(t *testing.T) {
	m := make(map[string]string)
	m["m01"] = "m01"
	m["m02"] = "m02"
	m["m03"] = "m03"
	fmt.Println(m)
	delete(m, "m01")
	fmt.Println(m)
}

func TestMapFind(t *testing.T) {
	m := make(map[string]string)
	m["m01"] = "m01"
	m["m02"] = "m02"
	m["m03"] = "m03"
	s, ok := m["03"]
	if ok {
		fmt.Printf("值存在：%v", s)
	} else {
		fmt.Println("值不存在")
	}
}

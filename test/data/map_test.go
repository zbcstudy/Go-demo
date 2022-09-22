package data

import (
	"fmt"
	"sort"
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

// 对map进行排序
func TestMapSort(t *testing.T) {
	target := make(map[int]int, 10)
	target[10] = 100
	target[1] = 13
	target[4] = 56
	target[8] = 33

	fmt.Println(target)

	var keys []int
	for k, _ := range target {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	fmt.Println(keys)

	for _, key := range keys {
		v := target[key]
		fmt.Printf("key:%v,value:%v\n", key, v)
	}
}

//func TestMapWithUtil(t *testing.T)  {
//	sMap := make(maputil.SMap)
//	sMap["abc"] = "123"
//	sMap["zbc"] = "456"
//	spew.Dump(&sMap)
//	fmt.Println(sMap.Get("zbc"))
//}

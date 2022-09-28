package _interface

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

// Len is the number of elements in the collection.
func (h HeroSlice) Len() int {
	return len(h)
}

// 按从小到大排序
func (h HeroSlice) Less(i, j int) bool {
	return h[i].Age < h[j].Age
}

// Swap swaps the elements with indexes i and j.
func (h HeroSlice) Swap(i, j int) {
	var temp Hero
	temp = h[i]
	h[i] = h[j]
	h[j] = temp
}

func TestSortFunc(t *testing.T) {
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: "hero" + strconv.Itoa(i),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}

	fmt.Println("排序前字段：")
	for _, hero := range heros {
		fmt.Println(hero)
	}

	sort.Sort(heros)

	fmt.Println("排序后字段：")
	for _, hero := range heros {
		fmt.Println(hero)
	}

}

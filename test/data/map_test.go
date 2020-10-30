package data

import "testing"

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

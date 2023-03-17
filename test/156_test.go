package main

import (
	"fmt"
	"net/url"
	"testing"
)

type Query struct {
	// type Values map[string][]string
	url.Values
}

func TestExtend(t *testing.T) {
	q := Query{}
	q.Values["name"] = []string{"test extend"}
	fmt.Println(q.Get("name"))
}

func TestExtend2(t *testing.T) {
	q := Query{
		Values: make(map[string][]string),
	}
	fmt.Println("len:", len(q.Values))
	q.Values["name"] = []string{"test extend"}
	fmt.Println(q.Get("name"))
}

package main

import (
	"fmt"
	"net/url"
	"testing"
)

type Query struct {
	url.Values
}

func TestExtend(t *testing.T) {
	q := Query{}
	q.Values["name"] = []string{"test extend"}
	fmt.Println(q.Get("name"))
}

package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	src := "bind 0.0.0.0\n" +
		"port 6399\n" +
		"appendonly yes\n" +
		"peers a,b"
	scanner := bufio.NewScanner(strings.NewReader(src))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

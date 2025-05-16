package main

import (
	"fmt"
	"os/exec"
	"os/user"
	"testing"
)

func TestExec(t *testing.T) {
	cmd := exec.Command("go", "version")
	output, _ := cmd.Output()
	fmt.Println(string(output))
	fmt.Println(cmd.Process.Pid)

	User, _ := user.Current()
	fmt.Println(User.HomeDir)
}

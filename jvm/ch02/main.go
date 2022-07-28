package main

import (
	"fmt"
)

// go install ./
func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJvm(cmd)
	}
}

func startJvm(cmd *Cmd) {
	fmt.Printf("classpath: %s class: %s args: %v\n", cmd.cpOption, cmd.class, cmd.args)
}

package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "地址")
	flag.IntVar(&port, "p", 8080, "端口")

	flag.Parse()

	fmt.Println("user:", user, "pwd:", pwd, "host:", host, "port:", port)
}

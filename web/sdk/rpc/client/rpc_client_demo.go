package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "10.0.1.20:2234")
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	var reply string
	err = client.Call("helloService.Hello", "name-zbc", &reply)
	if err != nil {
		log.Fatal("client call error:", err)
	}
	fmt.Println(reply)
}

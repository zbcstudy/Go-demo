package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceClient struct {
	*rpc.Client
}

//var HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network string, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (h *HelloServiceClient) Hello(request string, reply *string) error {
	return h.Client.Call(HelloServiceName+".Hello", request, reply)
}

func main() {
	client, err := DialHelloService("tcp", "10.0.1.20:2234")

	if err != nil {
		log.Fatal("rpc dial err:", err)
	}

	var reply string
	err = client.Hello("name-ert", &reply)
	if err != nil {
		log.Fatal("call err:", err)
	}

	fmt.Println(reply)
}

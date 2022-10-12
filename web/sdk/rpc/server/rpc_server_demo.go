package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello: " + request
	return nil
}

func main() {
	_ = rpc.RegisterName("helloService", new(HelloService))

	listen, err := net.Listen("tcp", ":2234")
	if err != nil {
		log.Fatal("listen tcp err:", err)
	}

	conn, err := listen.Accept()
	if err != nil {
		log.Fatal("accept error:", err)
	}

	rpc.ServeConn(conn)

}

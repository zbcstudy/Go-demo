package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func Echo(ws *websocket.Conn) {

	for {
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("can't receive")
			break
		}

		fmt.Println("receive back from client:", reply)

		msg := "receive:" + reply

		fmt.Println("send to client:", msg)

		if err := websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("can't send")
			break
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}

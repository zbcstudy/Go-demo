package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "welcome to the home page 111")
	})
	n := negroni.Classic()
	// go自带的路由
	n.UseHandler(serveMux)
	//n.UseHandler(muxRouter())
	_ = http.ListenAndServe(":3001", n)
}

// http://127.0.0.1:3001/first
func muxRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/first", firstHandler)
	return r
}

func firstHandler(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write([]byte("first route handler"))
}

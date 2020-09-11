package main

import (
	"fmt"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "welcome to the home page 111")
	})
	n := negroni.Classic()
	n.UseHandler(mux)
	_ = http.ListenAndServe(":3000", n)
}

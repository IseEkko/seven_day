package main

import (
	"fmt"
	"log"
	"net/http"
)

type engine struct {
}

func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Println("lzz")
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
func main() {
	en := new(engine)
	log.Fatal(http.ListenAndServe(":9999", en))
}

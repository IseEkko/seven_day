package main

import (
	gee "7Days/Gee"
	"fmt"
	"net/http"
)

func main() {
	en := gee.New()
	en.POST("/", Index)
	en.Run(":8080")
}

func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello world")
}

package main

import (
	"fmt"
	"net/http"
)

const (
	// all interfaces and port number
	helloAddress = ":2020"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello world")
	fmt.Fprintf(w, "hello world\n")
}

func main() {
	http.HandleFunc("/hello", hello)

	fmt.Printf("Opening HTTP hello-world server at '%v'\n", helloAddress)
	err := http.ListenAndServe(helloAddress, nil)
	if err != nil {
		fmt.Printf("HTTP server failed: %v\n", err)
	}
}

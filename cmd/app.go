package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	message := "This is a simple message"
	w.Write([]byte(message))
}

func main() {
	fmt.Println("Starting golang server...")
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":80", nil)
}

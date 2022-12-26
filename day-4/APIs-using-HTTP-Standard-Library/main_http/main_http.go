package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloWorldHandler)
	http.ListenAndServe(":8080", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

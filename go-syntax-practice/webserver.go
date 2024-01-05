package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Your server is live on port :8080, visit http://localhost:300/name to test server")
	http.HandleFunc("/", helloServer)
	http.ListenAndServe(":8080", nil)
}

func helloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", req.URL.Path[1:])
}
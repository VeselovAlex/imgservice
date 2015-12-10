package main

import (
	"net/http"
	"fmt"
)

func handleHome(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "<h1>Hello, world</h1>")
}

func main() {
	http.HandleFunc("/", handleHome)
	http.ListenAndServe(":8080", nil)
}
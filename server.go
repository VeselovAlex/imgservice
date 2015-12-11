package main

import (
	"net/http"

	"github.com/VeselovAlex/imgservice/handlers"
)

func main() {
	http.Handle("/", handlers.Home)
	http.ListenAndServe(":8080", nil)
}

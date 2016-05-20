package main

import (
	"fmt"
	"net/http"

	"github.com/aubm/github-aggregator/api"
)

func main() {
	http.HandleFunc("/hello", api.Hello)
	fmt.Println("Application started on port 8080")
	http.ListenAndServe(":8080", nil)
}

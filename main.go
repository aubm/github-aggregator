package main

import (
	"fmt"
	"net/http"

	"github.com/aubm/github-aggregator/api"
	"github.com/aubm/github-aggregator/github"
)

func main() {
	dir := "/Users/aurelienbaumann/Desktop/clones"
	reposManager := github.ReposManager{Dir: dir}
	handlers := api.ReposHandlers{Manager: reposManager}

	http.HandleFunc("/", handlers.CloneRepos)
	fmt.Println("Application started on port 8080")
	http.ListenAndServe(":8080", nil)
}

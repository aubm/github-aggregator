package api

import (
	"fmt"
	"net/http"

	"github.com/aubm/github-aggregator/github"
)

type listCloner interface {
	List(user string) ([]github.Repo, error)
	Clone(repo github.Repo) error
}

type ReposHandlers struct {
	Manager listCloner
}

func (h ReposHandlers) CloneRepos(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")

	repos, err := h.Manager.List(user)
	if err != nil {
		http.Error(w, "Invalid user", http.StatusBadRequest)
		return
	}
	done := make(chan (bool), len(repos))
	for _, repo := range repos {
		go func(repo github.Repo) {
			err := h.Manager.Clone(repo)
			if err == nil {
				fmt.Fprintf(w, "%v successfully cloned\n", repo.Name)
			} else {
				fmt.Fprintf(w, "failed to clone %v\n", repo.Name)
			}
			done <- true
		}(repo)
	}
	for range repos {
		<-done
	}

	w.Write([]byte("Hello, " + user))
}

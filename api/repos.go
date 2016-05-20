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
	for _, repo := range repos {
		err := h.Manager.Clone(repo)
		if err == nil {
			fmt.Fprintf(w, "%v successfully cloned\n", repo.Name)
		} else {
			fmt.Fprintf(w, "failed to clone %v\n", repo.Name)
		}
	}

	w.Write([]byte("Hello, " + user))
}

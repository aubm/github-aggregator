package api

import (
	"net/http"

	"github.com/aubm/github-aggregator/github"
)

type listCloner interface {
	List(user string) []github.Repo
	Clone(repo github.Repo)
}

type ReposHandlers struct {
	Manager listCloner
}

func (h ReposHandlers) CloneRepos(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")

	repos := h.Manager.List(user)
	for _, repo := range repos {
		h.Manager.Clone(repo)
	}

	w.Write([]byte("Hello, " + user))
}

package api

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/aubm/github-aggregator/github"
)

type mockListCloner struct {
	LastListArg  string
	NbCloneCalls int32
}

func (lc *mockListCloner) List(user string) ([]github.Repo, error) {
	lc.LastListArg = user
	return []github.Repo{{URL: "foo"}, {URL: "bar"}}, nil
}

func (lc *mockListCloner) Clone(github.Repo) error {
	lc.NbCloneCalls = atomic.AddInt32(&lc.NbCloneCalls, 1)
	return nil
}

func TestCloneRepos(t *testing.T) {
	// Given
	manager := &mockListCloner{}
	handler := ReposHandlers{Manager: manager}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/?user=aubm", nil)

	// When
	handler.CloneRepos(w, r)

	// Then
	if manager.LastListArg != "aubm" {
		t.Errorf("last list arg is %v, expected aubm", manager.LastListArg)
	}
	if manager.NbCloneCalls != 2 {
		t.Errorf("nb clone calls is %v, expected 2", manager.NbCloneCalls)
	}
}

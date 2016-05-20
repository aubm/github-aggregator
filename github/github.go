package github

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"
)

type Repo struct {
	Name string `json:"name"`
	URL  string `json:"ssh_url"`
}

type ReposManager struct {
	Dir string
}

func (m ReposManager) List(user string) ([]Repo, error) {
	resp, err := http.Get("https://api.github.com/users/" + user + "/repos")
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	repos := []Repo{}
	if err := json.Unmarshal(body, &repos); err != nil {
		return nil, err
	}
	return repos, nil
}

func (m ReposManager) Clone(repo Repo) error {
	cmd := exec.Command("git", "clone", repo.URL)
	cmd.Dir = m.Dir
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

package github

type Repo struct {
	Name string
	URL  string
}

type ReposManager struct{}

func (m ReposManager) List(user string) []Repo {

}

func (m ReposManager) Clone(repo Repo) {

}

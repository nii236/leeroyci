// GitHub provides all structs to unmarshal a GitHub webhook.
package callbacks

import (
	"encoding/json"
	"ironman/logging"
	"log"
	"strings"
	"time"
)

type GitHubCallback struct {
	Ref         string
	After       string
	Before      string
	Created     bool
	Deleted     bool
	Forced      bool
	Compare     string
	Commits     []Commit
	Head_commit Commit
	Repository  Repository
	Pusher      GitUser
}

type Commit struct {
	Id        string
	Distinct  bool
	Message   string
	Timestamp string
	Url       string
	Author    GitHubUser
	Committer GitHubUser
	Added     []string
	Removed   []string
	Modified  []string
}

type GitHubUser struct {
	Name     string
	Email    string
	Username string
}

type Repository struct {
	Id            int64
	Name          string
	Url           string
	Description   string
	Watchers      int
	Stargazers    int
	Forks         int
	Size          int
	Owner         GitUser
	Private       bool
	Open_issues   int
	Has_issues    bool
	Has_downloads bool
	Has_wiki      bool
	Created_at    int64
	Pushed_at     int64
	Master_branch string
}

type GitUser struct {
	Name  string
	Email string
}

// Branch returns the name of the branch.
func (g *GitHubCallback) Branch() string {
	s := strings.Split(g.Ref, "/")
	return s[2]
}

// URL returns the URL for the repository
func (g *GitHubCallback) URL() string {
	return g.Repository.Url
}

// By returns who pushed / triggered the callback. Format Name <email>.
func (g *GitHubCallback) By() (string, string) {
	return g.Pusher.Name, g.Pusher.Email
}

// Returns if this commit should be build. Do not build if the branch was
// deleted for example.
func (g *GitHubCallback) ShouldBuild() bool {
	if g.Deleted == true {
		return false
	}
	return true
}

// Returns the ID of the head commit.
func (g *GitHubCallback) Commit() string {
	return g.Head_commit.Id
}

// Parse a GitHub request body and add it to the build queue.
func parseGitHub(jobs chan logging.Job, body []byte) {
	var cb GitHubCallback
	err := json.Unmarshal(body, &cb)

	name, email := cb.By()

	j := logging.Job{
		URL:       cb.URL(),
		Branch:    cb.Branch(),
		Timestamp: time.Now(),
		Commit:    cb.Commit(),
		Name:      name,
		Email:     email,
	}

	if err != nil {
		log.Println(string(body))
		panic("Could not unmarshal request")
	}

	if cb.ShouldBuild() == true {
		jobs <- j
	} else {
		log.Println("Not adding", cb.URL(), cb.Branch(), "to build queue")
	}
}

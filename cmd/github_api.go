package cmd

import (
	"github.com/google/go-github/github"
	"context"
	"strings"
)

// GithubAPI type which holds the Repository name and the owner name
type GithubAPI struct {
	RepoName string
	Owner string
}

// NewGithubAPI instantiates new GithubAPI type given the repoPath
func NewGithubAPI(repoPath string) GithubAPI {
	repoName, owner := parseRepoName(repoPath)
	return GithubAPI{
		RepoName: repoName,
		Owner: owner,
	}
}

func parseRepoName(repoPath string) (string, string) {
	values := strings.Split(repoPath, "/")
	owner := values[0]
	repoName := values[1]
	return repoName, owner
}

// FetchCommits - Makes call to Github service and fetches the Commit Activity for the repo
func (githubAPI GithubAPI) FetchCommits() ([]*github.WeeklyCommitActivity, error){
	ctx := context.Background()
	client := github.NewClient(nil)
	commitActivities, _, err := client.Repositories.ListCommitActivity(ctx, githubAPI.Owner, githubAPI.RepoName)
	if err != nil {
		return nil, err
	}
	return commitActivities, nil
}
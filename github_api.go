package main

import (
	"github.com/google/go-github/github"
	"context"
	"strings"
)

type GithubAPI struct {
	RepoName string
	Owner string
}

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

func (githubAPI GithubAPI) FetchCommits() []*github.WeeklyCommitActivity{
	ctx := context.Background()
	client := github.NewClient(nil)
	commitActivities, _, err := client.Repositories.ListCommitActivity(ctx, githubAPI.Owner, githubAPI.RepoName)
	if err != nil {
		return nil
	}
	return commitActivities
}
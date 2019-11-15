package main

import (
	"github.com/google/go-github/github"
	"context"
)

type GithubAPI struct {
	baseUrl string	
}

func NewGithubAPI(baseUrl string) GithubAPI {
	return GithubAPI{
		baseUrl: baseUrl,
	}
}

func (githubAPI GithubAPI) FetchCommits(owner string, repoName string) []*github.WeeklyCommitActivity{
	ctx := context.Background()
	client := github.NewClient(nil)
	commitActivities, _, err := client.Repositories.ListCommitActivity(ctx, owner, repoName)
	if err != nil {
		return nil
	}
	return commitActivities
}
package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/worldofprasanna/github-stats"
)

func TestFetchCommits(t *testing.T) {

	t.Run("should split the repo path properly", func(t *testing.T) {
		githubAPI := main.NewGithubAPI("google/kubernetes")
		assert.Equal(t, githubAPI.Owner, "google", "should have the proper owner name")
		assert.Equal(t, githubAPI.RepoName, "kubernetes", "should have the proper repo name")
	})

	t.Run("should fetch the tweets properly", func(t *testing.T) {
		githubAPI := main.NewGithubAPI("kubernetes/kubernetes")
		values := githubAPI.FetchCommits()
		assert.Equal(t, len(values), 52, "should fetch 52 weeks record")
	})
}

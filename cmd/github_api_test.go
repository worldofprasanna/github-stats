package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/worldofprasanna/github-stats/cmd"
)

func TestFetchCommits(t *testing.T) {

	t.Run("should split the repo path properly", func(t *testing.T) {
		githubAPI := cmd.NewGithubAPI("google/kubernetes")
		assert.Equal(t, githubAPI.Owner, "google", "should have the proper owner name")
		assert.Equal(t, githubAPI.RepoName, "kubernetes", "should have the proper repo name")
	})

	t.Run("should fetch the tweets properly", func(t *testing.T) {
		githubAPI := cmd.NewGithubAPI("kubernetes/kubernetes")
		values, err := githubAPI.FetchCommits()
		assert.Equal(t, len(values), 52, "should fetch 52 weeks record")
		assert.Equal(t, err, nil, "should not have errored out")
	})
}

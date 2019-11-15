package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/worldofprasanna/github-stats"
)

func TestFetchCommits(t *testing.T) {
	t.Run("should fetch the tweets properly", func(t *testing.T) {
		githubAPI := main.NewGithubAPI("https://api.github.com/")
		values := githubAPI.FetchCommits("kubernetes", "kubernetes")
		assert.Equal(t, len(values), 52, "should fetch 52 weeks record")
	})
}

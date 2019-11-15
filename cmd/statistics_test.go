package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/google/go-github/github"
	"gitlab.com/worldofprasanna/github-stats"
)

func TestStatistics(t *testing.T) {

	t.Run("should fetch the proper data from Github WeeklyCommitActivity", func(t *testing.T) {
		weeklyCommitActivity := github.WeeklyCommitActivity{
			Days: []int{ 10, 11, 12, 13, 14, 15, 16 },
		}
		expectedData := make(map[string]int)
		expectedData["sunday"] = 10
		expectedData["monday"] = 11
		expectedData["tuesday"] = 12
		expectedData["wednesday"] = 13
		expectedData["thursday"] = 14
		expectedData["friday"] = 15
		expectedData["saturday"] = 16
		actualData := main.ParseCommitActivity(&weeklyCommitActivity)
		assert.Equal(t, actualData, expectedData, "should have parsed the data properly")
	})

	t.Run("should aggregate the weekly commit activity", func(t *testing.T) {
		commit1 := main.ParseCommitActivity(&github.WeeklyCommitActivity{
			Days: []int{ 10, 11, 12, 13, 14, 15, 16 },
		})
		commit2 := main.ParseCommitActivity(&github.WeeklyCommitActivity{
			Days: []int{ 10, 11, 12, 13, 14, 15, 16 },
		})
		commit3 := main.ParseCommitActivity(&github.WeeklyCommitActivity{
			Days: []int{ 10, 11, 12, 13, 14, 15, 16 },
		})
		expectedData := make(map[string]int)
		expectedData["sunday"] = 30
		expectedData["monday"] = 33
		expectedData["tuesday"] = 36
		expectedData["wednesday"] = 39
		expectedData["thursday"] = 42
		expectedData["friday"] = 45
		expectedData["saturday"] = 48
		actualData := main.AggregateCommitActivities([]map[string]int{commit1, commit2, commit3})
		assert.Equal(t, actualData, expectedData, "should have aggregated the data properly")
	})

	t.Run("should find most commits of the day", func(t *testing.T) {
		commit := make(map[string]int)
		commit["sunday"] = 30
		commit["monday"] = 33
		commit["tuesday"] = 36
		commit["wednesday"] = 49
		commit["thursday"] = 42
		commit["friday"] = 45
		commit["saturday"] = 48
		maxDay, maxCommit := main.FindMostCommitsDay(commit)
		assert.Equal(t, maxDay, "wednesday", "should have found the max day properly")
		assert.Equal(t, maxCommit, 49, "should have found the max commit properly")
	})

	// This is a flaky test, as it depends on the actual Github API and the result may change in future.
	t.Run("should find the max number of commit and day for the repository", func(t *testing.T) {
		statistics := main.NewStatistics("kubernetes/kubernetes", 52, "asc")
		result := statistics.ActiveDayInRepo()
		assert.Equal(t, result, "wednesday 28", "should have found the max commit and day for it")
	})
}

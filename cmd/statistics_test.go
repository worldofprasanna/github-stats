package cmd_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/google/go-github/github"
	"github.com/worldofprasanna/github-stats/cmd"
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
		actualData := cmd.ParseCommitActivity(&weeklyCommitActivity)
		assert.Equal(t, actualData, expectedData, "should have parsed the data properly")
	})

	t.Run("should aggregate the weekly commit activity", func(t *testing.T) {
		commit1 := cmd.ParseCommitActivity(&github.WeeklyCommitActivity{
			Days: []int{ 10, 11, 12, 13, 14, 15, 16 },
		})
		commit2 := cmd.ParseCommitActivity(&github.WeeklyCommitActivity{
			Days: []int{ 10, 11, 12, 13, 14, 15, 16 },
		})
		commit3 := cmd.ParseCommitActivity(&github.WeeklyCommitActivity{
			Days: []int{ 10, 11, 12, 13, 14, 15, 16 },
		})
		expectedData := make(map[string]int)
		expectedData["sunday"] = 10
		expectedData["monday"] = 11
		expectedData["tuesday"] = 12
		expectedData["wednesday"] = 13
		expectedData["thursday"] = 14
		expectedData["friday"] = 15
		expectedData["saturday"] = 16
		actualData := cmd.AggregateCommitActivities([]map[string]int{commit1, commit2, commit3}, 3)
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
		maxDay, maxCommit := cmd.FindMostCommitsDay(commit)
		assert.Equal(t, maxDay, "wednesday", "should have found the max day properly")
		assert.Equal(t, maxCommit, 49, "should have found the max commit properly")
	})

	t.Run("should filter the last n commits", func(t *testing.T) {
		c1Total := 100
		c2Total := 101
		c3Total := 102
		c4Total := 103
		c5Total := 104
		c6Total := 105
		commit1 := &github.WeeklyCommitActivity{
			Total: &c1Total,
		}
		commit2 := &github.WeeklyCommitActivity{
			Total: &c2Total,
		}
		commit3 := &github.WeeklyCommitActivity{
			Total: &c3Total,
		}
		commit4 := &github.WeeklyCommitActivity{
			Total: &c4Total,
		}
		commit5 := &github.WeeklyCommitActivity{
			Total: &c5Total,
		}
		commit6 := &github.WeeklyCommitActivity{
			Total: &c6Total,
		}
		filteredCommits := cmd.FilterLastNRecords([]*github.WeeklyCommitActivity{
			commit1,
			commit2,
			commit3,
			commit4,
			commit5,
			commit6,
		}, 4)
		assert.Equal(t, len(filteredCommits), 4, "should have filtered the commits properly")
		for i, val := range filteredCommits {
			assert.Equal(t, (102+i), *val.Total, (102+i), "should have properly filtered the values")
		}
	})

	// This is a flaky test, as it depends on the actual Github API and the result may change in future.
	t.Run("should find the max number of commit and day for the repository", func(t *testing.T) {
		statistics, err := cmd.NewStatistics("kubernetes/kubernetes", 52, "asc")
		result := statistics.ActiveDayInRepo()
		assert.NotEqual(t, result, "", "should have found the max commit and day for it. Ignore the actual value as it would change in realtime.")
		assert.Equal(t, err, nil, "should not be errored out")
	})

	t.Run("should sort the commits based on commit count", func(t *testing.T) {
		commit := cmd.ParseCommitActivity(&github.WeeklyCommitActivity{
			Days: []int{ 10, 11, 12, 13, 14, 15, 16 },
		})
		actualCommits := cmd.SortCommitsPerDay(commit, "desc")
		for i, val := range actualCommits {
			assert.Equal(t, (16-i), val.Value, "should have properly sorted the values")
		}
	})
}

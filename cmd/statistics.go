package cmd

import (
	"fmt"
	"sort"
	"github.com/google/go-github/github"
)

type Statistics struct {
	weeks int
	sort string
	commitActivities []map[string]int
}

type SortedCommit struct {
	Key   string
	Value int
}

func NewStatistics(repoPath string, weeks int, sort string) *Statistics {
	githubAPI := NewGithubAPI(repoPath)
	rawCommitActivities := githubAPI.FetchCommits()
	lastCommitActivities := FilterLastNRecords(rawCommitActivities, weeks)
	commitActivities := make([]map[string]int, len(lastCommitActivities))
	for i, val := range lastCommitActivities {
			commitActivities[i] = ParseCommitActivity(val)
	}
	return &Statistics{
		weeks: weeks,
		commitActivities: commitActivities,
		sort: sort,
	}
}

func FilterLastNRecords(commits []*github.WeeklyCommitActivity, weeks int) []*github.WeeklyCommitActivity {
	var lastCommitActivities []*github.WeeklyCommitActivity
	maxValue := len(commits)
	if weeks < maxValue {
		fromWeeks := maxValue - weeks
		lastCommitActivities = commits[fromWeeks:maxValue]
	} else {
		lastCommitActivities = commits
	}
	return lastCommitActivities
}

func (s *Statistics) ActiveDayInRepo() string {
	aggregatedCommitActivities := AggregateCommitActivities(s.commitActivities, s.weeks)
	maxCommitDay, maxCommit := FindMostCommitsDay(aggregatedCommitActivities)
	return fmt.Sprintf("%s %d", maxCommitDay, maxCommit)
}

func AggregateCommitActivities(commits []map[string]int, totalWeeks int) map[string]int {
	aggregatedCommits := make(map[string]int)
	for _, commit := range commits {
		aggregatedCommits["sunday"] += commit["sunday"]
		aggregatedCommits["monday"] += commit["monday"]
		aggregatedCommits["tuesday"] += commit["tuesday"]
		aggregatedCommits["wednesday"] += commit["wednesday"]
		aggregatedCommits["thursday"] += commit["thursday"]
		aggregatedCommits["friday"] += commit["friday"]
		aggregatedCommits["saturday"] += commit["saturday"]
	}

	for day, count := range aggregatedCommits {
		aggregatedCommits[day] = count / totalWeeks
	}
	return aggregatedCommits
}

func FindMostCommitsDay(commits map[string]int) (string, int) {
	var maxDay string
	maxCommits := 0

	for day, noOfCommits := range commits {
		if maxCommits < noOfCommits {
			maxCommits = noOfCommits
			maxDay = day
		}
	}
	return maxDay, maxCommits
}

func ParseCommitActivity(ca *github.WeeklyCommitActivity) map[string]int {
	commitActivity := make(map[string]int)
	commitsPerDay := ca.Days
	commitActivity["sunday"] = commitsPerDay[0]
	commitActivity["monday"] = commitsPerDay[1]
	commitActivity["tuesday"] = commitsPerDay[2]
	commitActivity["wednesday"] = commitsPerDay[3]
	commitActivity["thursday"] = commitsPerDay[4]
	commitActivity["friday"] = commitsPerDay[5]
	commitActivity["saturday"] = commitsPerDay[6]
	
	return commitActivity
}

func (s *Statistics) AverageCommitPerDay() {
	commits := AggregateCommitActivities(s.commitActivities, s.weeks)
	sortedCommits := SortCommitsPerDay(commits, s.sort)
	for _, commit := range sortedCommits {
		fmt.Printf("%s has avg of %d commits\n", commit.Key, commit.Value)
	}
}

func SortCommitsPerDay(commits map[string]int, sortBy string) []SortedCommit {
	var sortedCommits []SortedCommit
  for k, v := range commits {
    sortedCommits = append(sortedCommits, SortedCommit{k, v})
	}
	if sortBy == "asc" {
		sort.Slice(sortedCommits, func(i, j int) bool {
			return sortedCommits[i].Value < sortedCommits[j].Value
		})
	} else {
		sort.Slice(sortedCommits, func(i, j int) bool {
			return sortedCommits[i].Value > sortedCommits[j].Value
		})
	}
	return sortedCommits
}
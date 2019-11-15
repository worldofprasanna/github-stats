package main

import (
	"fmt"
	"github.com/google/go-github/github"
)

type Statistics struct {
	weeks int
	order string
	commitActivities []map[string]int
}

func NewStatistics(repoPath string, weeks int, order string) *Statistics {
	githubAPI := NewGithubAPI(repoPath)
	rawCommitActivities := githubAPI.FetchCommits()
	commitActivities := make([]map[string]int, len(rawCommitActivities))
	for i, val := range rawCommitActivities {
			commitActivities[i] = ParseCommitActivity(val)
	}
	return &Statistics{
		weeks: weeks,
		order: order,
		commitActivities: commitActivities,
	}
}

func (s *Statistics) ActiveDayInRepo() string {
	aggregatedCommitActivities := AggregateCommitActivities(s.commitActivities)
	maxCommitDay, maxCommit := FindMostCommitsDay(aggregatedCommitActivities)
	return fmt.Sprintf("%s %d", maxCommitDay, maxCommit / s.weeks)
}

func AggregateCommitActivities(commits []map[string]int) map[string]int {
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
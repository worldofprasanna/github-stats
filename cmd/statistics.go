package cmd

import (
	"fmt"
	"sort"
	"github.com/google/go-github/github"
)

// Statistics type which holds all the business logic
type Statistics struct {
	weeks int
	sort string
	commitActivities []map[string]int
}

// SortedCommit type has key - day, value - no of commits.
// Use to store the commits based on no of commits order.
type SortedCommit struct {
	Key   string
	Value int
}

// NewStatistics function instantiates new Statistics based on the args
func NewStatistics(repoPath string, weeks int, sort string) (*Statistics, error) {
	githubAPI := NewGithubAPI(repoPath)
	rawCommitActivities, err := githubAPI.FetchCommits()
	if err != nil {
		return nil, err
	}
	lastCommitActivities := FilterLastNRecords(rawCommitActivities, weeks)
	commitActivities := make([]map[string]int, len(lastCommitActivities))
	for i, val := range lastCommitActivities {
			commitActivities[i] = ParseCommitActivity(val)
	}
	return &Statistics{
		weeks: weeks,
		commitActivities: commitActivities,
		sort: sort,
	}, nil
}

var daysOfWeek = []string {"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

// FilterLastNRecords - Used to filter the records for the last n weeks given
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

// ActiveDayInRepo - Returns the most no of commits made for the day of week
func (s *Statistics) ActiveDayInRepo() string {
	aggregatedCommitActivities := AggregateCommitActivities(s.commitActivities, s.weeks)
	maxCommitDay, maxCommit := FindMostCommitsDay(aggregatedCommitActivities)
	return fmt.Sprintf("%s %d", maxCommitDay, maxCommit)
}

// AggregateCommitActivities - Returns all the commit aggregated for each day
func AggregateCommitActivities(commits []map[string]int, totalWeeks int) map[string]int {
	aggregatedCommits := make(map[string]int)
	for _, commit := range commits {
		for _, day := range daysOfWeek {
			aggregatedCommits[day] += commit[day]
		}
	}

	for day, count := range aggregatedCommits {
		aggregatedCommits[day] = count / totalWeeks
	}
	return aggregatedCommits
}

// FindMostCommitsDay - Returns the most no of commits along with the day of the week
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

// ParseCommitActivity - Converts object WeeklyCommitActivity into a map of { day => no of commits }
func ParseCommitActivity(ca *github.WeeklyCommitActivity) map[string]int {
	commitActivity := make(map[string]int)
	commitsPerDay := ca.Days
	for i, day := range daysOfWeek {
		commitActivity[day] = commitsPerDay[i]
	}
	
	return commitActivity
}

// AverageCommitPerDay - Returns the average no of commits made for each day of the week
func (s *Statistics) AverageCommitPerDay() {
	commits := AggregateCommitActivities(s.commitActivities, s.weeks)
	sortedCommits := SortCommitsPerDay(commits, s.sort)
	for _, commit := range sortedCommits {
		fmt.Printf("%s has avg of %d commits\n", commit.Key, commit.Value)
	}
}

// SortCommitsPerDay - Sorts the commits based on the no of commits either in ascending or descending order
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
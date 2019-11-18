package cmd

import (
	"fmt"
	"strings"
	"errors"
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

var listAverageCommitsCmd = &cobra.Command{
	Use:   "listAverageCommits",
	Short: "list the commit activity for the repository",
	Long: `params: repo name (string), sort (asc|desc)

Eg: ./app listAverageCommits --sort=desc kubernetes/kubernetes`,
	RunE: func(cmd *cobra.Command, args []string) error {
    sort, err := cmd.Flags().GetString("sort")
    if err != nil {
      return err
		}
		if sort != "asc" && sort != "desc" {
			return errors.New("Sort paramenter should be either asc / desc")
		}
    repoName := strings.Join(args, "")
    if !strings.Contains(repoName, "/") {
      return errors.New("Not a valid repo path. Specify in format <owner/repo>")
    }
    fmt.Printf("Going to display metrics for [repo - %s, sort order - %s]\n", repoName, sort)
	  statistics, err := NewStatistics(repoName, 52, sort)
	  if err != nil {
		  log.Error("Error occurred while fetching the commits from Github: ", err.Error())
		  return nil
	  }
    fmt.Println("Result:")
    statistics.AverageCommitPerDay()
    return nil
	},
}

func init() {
	listAverageCommitsCmd.Flags().StringP("sort", "s", "asc", "sort the list in asc / desc order")
	rootCmd.AddCommand(listAverageCommitsCmd)
}

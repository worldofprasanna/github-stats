package cmd

import (
  "fmt"
  "os"
  "strings"
  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "github-stats",
  Short: "Get commit info for the Github Repository",
  Long: `Given a repository name, fetch the stats based on the commit history.
  Please check help for more detailed instructions.`,
  RunE: func(cmd *cobra.Command, args []string) error {
    weeks, err := cmd.Flags().GetInt("weeks")
    if err != nil {
      return err
    }
    sort, err := cmd.Flags().GetString("sort")
    if err != nil {
      return err
    }
    repoName := strings.Join(args, "")
    fmt.Printf("Going to collect metrics for [repo - %s, weeks - %d, sorted by - %s]\n", repoName, weeks, sort)
    statistics := NewStatistics(repoName, weeks, sort)
    result := statistics.ActiveDayInRepo()
    fmt.Printf("Result: %s\n", result)
    return nil
  },
}

var weeks int
var sort string

func init() {
  rootCmd.Flags().IntVarP(&weeks, "weeks", "w", 52, "No of weeks to process")
  rootCmd.Flags().StringVarP(&sort, "sort", "s", "asc", "Sort the statistics [asc|desc]")
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

package cmd

import (
  "fmt"
  "os"
  "strings"
  "errors"
  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "github-stats",
  Short: "Get commit info for the Github Repository",
  Long: `Given a repository name, fetch the stats based on the commit history.
  Eg: github-stats kubernetes/kubernetes --weeks=20
  Please check help for more detailed instructions.`,
  RunE: func(cmd *cobra.Command, args []string) error {
    weeks, err := cmd.Flags().GetInt("weeks")
    if err != nil {
      return err
    }
    repoName := strings.Join(args, "")
    if !strings.Contains(repoName, "/") {
      return errors.New("Not a valid repo path. Specify in format <owner/repo>")
    }
    fmt.Printf("Going to collect metrics for [repo - %s, weeks - %d]\n", repoName, weeks)
    statistics := NewStatistics(repoName, weeks)
    result := statistics.ActiveDayInRepo()
    fmt.Printf("Result: %s\n", result)
    return nil
  },
}

var weeks int

func init() {
  rootCmd.Flags().IntVarP(&weeks, "weeks", "w", 52, "No of weeks to process")
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

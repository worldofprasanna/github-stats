package cmd

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "github-stats",
  Short: "Get commit info for the Github Repository",
  Long: `Given a repository name, fetch the stats based on the commit history.

  There are 2 commands you can specify. 
    1. activeDay
    2. listAverageCommits
  Look at the corresponding help command to know more about it.

Eg: github-stats listAverageCommits --help`,
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

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
  Please check help for more detailed instructions.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hello World Cobra")
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

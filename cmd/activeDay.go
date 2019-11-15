package cmd

import (
	"fmt"
	"strings"
	"errors"

	"github.com/spf13/cobra"
)

var activeDayCmd = &cobra.Command{
	Use:   "activeDay",
	Short: "Fetch the active day of the week along with the average commit count",
	Long: `Use this command to fetch the most active day for the specified no of weeks.

Eg: github-stats activeDay --weeks=20 kubernetes/kubernetes

Please check help for more detailed instructions.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		weeks, err := cmd.Flags().GetInt("weeks")
		if err != nil {
			return err
		}
		if weeks < 1 || weeks > 52 {
			return errors.New("Not a valid week. Week should be in the range of 1 - 52")
		}
		repoName := strings.Join(args, "")
		if !strings.Contains(repoName, "/") {
			return errors.New("Not a valid repo path. Specify in format <owner/repo>")
		}
		fmt.Printf("Going to collect metrics for [repo - %s, weeks - %d]\n", repoName, weeks)
		statistics := NewStatistics(repoName, weeks, "")
		result := statistics.ActiveDayInRepo()
		fmt.Printf("Result: %s\n", result)
		return nil
	},
}

var weeks int

func init() {
  activeDayCmd.Flags().IntVarP(&weeks, "weeks", "w", 52, "No of weeks to process")
	rootCmd.AddCommand(activeDayCmd)
}
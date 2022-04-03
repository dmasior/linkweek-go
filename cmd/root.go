package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "linkweek",
		Short: "LinkWeek App",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

package cmd

/*
Copyright © 2022 Evans Owamoyo <evans.dev99@gmail.com>
*/

import (
	"fmt"
	"github.com/spf13/cobra"
)

// linesCmd represents the lines command
var linesCmd = &cobra.Command{
	Use:   "lines",
	Short: "returns the number of lines in all files of a directory",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(workersCount, isVerbose, args[0])
	},
}

func init() {
	rootCmd.AddCommand(linesCmd)
}

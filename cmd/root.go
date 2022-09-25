package cmd

/*
Copyright © 2022 Evans Owamoyo <evans.dev99@gmail.com>

*/

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "gomeasure",
		Short: "measure lines of code in a project",
		Long:  `gomeasure is a CLI tool that measures lines of code and number of files in a directory`,
	}
	workersCount int
	isVerbose    bool
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.PersistentFlags().BoolVarP(&isVerbose, "verbose", "v", false, "displays the files being processed and their line count when `true`")
	rootCmd.Flags().IntVarP(&workersCount, "workers", "w", 5, "number of workers that scan files concurrently")
}

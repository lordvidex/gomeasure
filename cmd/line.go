package cmd

/*
Copyright © 2022 Evans Owamoyo <evans.dev99@gmail.com>
*/

import (
	"errors"
	"fmt"
	"github.com/lordvidex/gomeasure/pkg"
	"github.com/spf13/cobra"
)

// linesCmd represents the lines command
var linesCmd = &cobra.Command{
	Use:   "line <directory>",
	Short: "returns the number of lines in all files of a directory",

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cobra.CheckErr(errors.New("<directory> argument is required"))
			return
		}
		cobra.CheckErr(processLines(args[0]))
	},
}

var countEmptyLines bool

func processLines(directory string) error {
	runner := &pkg.Runner{
		IncludedFiles:    include,
		ExcludedFiles:    exclude,
		ShouldCountEmpty: countEmptyLines,
		WorkersCount:     workersCount,
		Directory:        directory,
		ShouldCountLines: true,
	}
	results, err := runner.Run()
	var total int64 = 0
	for _, file := range results {
		if isVerbose {
			fmt.Printf("%30s: %d\n", file.FilePath, file.LinesCount)
		}
		total += file.LinesCount
	}
	fmt.Printf("\n%s has %d lines of code\n", directory, total)
	return err
}

func init() {
	rootCmd.AddCommand(linesCmd)
	linesCmd.Flags().BoolVarP(&countEmptyLines, "empty", "e", false, "add this flag to count empty lines, default value `false`")
	linesCmd.Flags().IntVarP(&workersCount, "workers", "w", 5, "number of workers that scan files concurrently")
}

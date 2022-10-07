/*
Copyright © 2022 Evans Owamoyo
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/lordvidex/gomeasure/pkg/gomeasure"
	"github.com/spf13/cobra"
)

// characterCmd represents the character command
var characterCmd = &cobra.Command{
	Use:   "character <directory|file>",
	Short: "processes the number of characters contained in the files in a directory",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cobra.CheckErr(errors.New("an argument is required"))
			return
		}
		cobra.CheckErr(processCharacters(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(characterCmd)
}

func processCharacters(file string) error {
	runner := &gomeasure.Runner{
		IncludedFiles:    include,
		ExcludedFiles:    exclude,
		ShouldCountEmpty: countEmptyLines,
		WorkersCount:     workersCount,
		Directory:        file,
		Action:           gomeasure.MeasureCharacter,
	}
	results, err := runner.Run()
	if err != nil {
		return err
	}

	var total int64 = 0
	for _, result := range results {
		if isVerbose {
			fmt.Printf("%30s -> %d characters\n", result.FilePath, result.Count)
		}
		total += result.Count
	}
	fmt.Printf("\n%s has %d total characters\n", file, total)
	return err
}

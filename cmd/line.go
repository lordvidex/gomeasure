package cmd

/*
Copyright © 2022 Evans Owamoyo <evans.dev99@gmail.com>
*/

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lordvidex/gomeasure/pkg/gomeasure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var lineConfig = gomeasure.NewConfig()

// linesCmd represents the lines command
var linesCmd = &cobra.Command{
	Use:   "line <directory>",
	Short: "returns the number of lines in all files of a directory",

	Args: cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, _ []string) {
		if !initLineConfig() {
			x := *generalConfig
			lineConfig = &x
		}

		// parse flags
		parseFlags(cmd, lineConfig)
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cobra.CheckErr(errors.New("<directory> argument is required"))
			return
		}
		cobra.CheckErr(processLines(args[0]))
	},
}

func processLines(directory string) error {
	runner := &gomeasure.Runner{
		Config:    lineConfig,
		Directory: directory,
		Action:    gomeasure.MeasureLine,
	}
	results, err := runner.Run()
	var total int64 = 0
	for _, file := range results {
		if lineConfig.IsVerbose {
			fmt.Printf("%30s -> %d lines \n", file.FilePath, file.Count)
		}
		total += file.Count
	}
	fmt.Printf("\n%s has %d lines of code\n", directory, total)
	return err
}

// initLineConfig initializes the lineConfig variable
// and returns true for a successful operation otherwise false
func initLineConfig() bool {
	data := viper.GetStringMap("line")
	if len(data) == 0 {
		return false
	}
	bytes, err := json.Marshal(data)
	cobra.CheckErr(err)
	cobra.CheckErr(json.Unmarshal(bytes, &lineConfig))
	return true
}

func init() {
	rootCmd.AddCommand(linesCmd)
	linesCmd.Flags().BoolP("empty", "e", false, "add this flag to count empty lines, default value `false`")
	linesCmd.Flags().IntP("workers", "w", 5, "number of workers that scan files concurrently")
}

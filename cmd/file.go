package cmd

/*
Copyright © 2022 Evans Owamoyo <evans.dev99@gmail.com>
*/
import (
	"errors"
	"fmt"
	"github.com/lordvidex/gomeasure/pkg"
	"github.com/spf13/cobra"
	"path/filepath"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use: "file",

	Short: "processes the number of files in a directory",
	Long:  `gomeasure file processes and returns the number of files in a directory / project.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cobra.CheckErr(errors.New("<directory> argument is required"))
			return
		}
		cobra.CheckErr(processFiles(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
}

func processFiles(directory string) error {
	runner := &pkg.Runner{
		IncludedFiles:    include,
		ExcludedFiles:    exclude,
		Directory:        directory,
		ShouldCountLines: false,
	}
	results, err := runner.Run()
	if err != nil {
		return err
	}
	if isVerbose {
		abs, err := filepath.Abs(directory)
		if err != nil {
			abs = directory // sorry lol :(
		}
		fmt.Println("Files contained in \"", abs, "\"")
		for _, file := range results {
			fmt.Println(file.FilePath)
		}
	}
	fmt.Printf("\n%s has %d files\n", directory, len(results))
	return nil
}

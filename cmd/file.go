package cmd

/*
Copyright © 2022 Evans Owamoyo <evans.dev99@gmail.com>
*/
import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/lordvidex/gomeasure/pkg/gomeasure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var fileConfig = gomeasure.NewConfig()

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use: "file <directory>",

	Short: "processes the number of files in a directory",
	Long:  `gomeasure file processes and returns the number of files in a directory / project.`,
	PreRun: func(cmd *cobra.Command, _ []string) {
		
		if !initFileConfig() {
			x := *generalConfig
			fileConfig = &x
		}
		parseFlags(cmd, fileConfig)
	},
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

func initFileConfig() bool {
	data := viper.GetStringMap("file")
	if len(data) == 0 {
		return false
	}
	bytes, err := json.Marshal(data)
	cobra.CheckErr(err)
	cobra.CheckErr(json.Unmarshal(bytes, &fileConfig))
	return true
}

func processFiles(directory string) error {
	runner := &gomeasure.Runner{
		Config:    fileConfig,
		Directory: directory,
		Action:    gomeasure.MeasureFile,
	}
	results, err := runner.Run()
	if err != nil {
		return err
	}
	if fileConfig.IsVerbose {
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

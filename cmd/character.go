/*
Copyright © 2022 Evans Owamoyo
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lordvidex/gomeasure/pkg/gomeasure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var characterConfig = gomeasure.NewConfig()

// characterCmd represents the character command
var characterCmd = &cobra.Command{
	Use:   "character <directory|file>",
	Short: "processes the number of characters contained in the files in a directory",
	PreRun: func(cmd *cobra.Command, args []string) {
		
		if !initCharacterConfig() {
			x := *generalConfig
			characterConfig = &x
		}
		// parse flags
		parseFlags(cmd, characterConfig)
	},
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
	characterCmd.Flags().IntP("workers", "w", 5, "number of concurrent workers to use")
}

// initCharacterConfig fetches the character section from the config
// files and sets the characterConfig variable
// returns true if the character section was found in the config files
// otherwise returns false
func initCharacterConfig() bool {
	data := viper.GetStringMap("character")
	if len(data) == 0 {
		return false
	}
	bytes, err := json.Marshal(data)
	cobra.CheckErr(err)
	cobra.CheckErr(json.Unmarshal(bytes, &characterConfig))
	return true
}

func processCharacters(file string) error {
	runner := &gomeasure.Runner{
		Config:    characterConfig,
		Directory: file,
		Action:    gomeasure.MeasureCharacter,
	}
	results, err := runner.Run()
	if err != nil {
		return err
	}

	var total int64 = 0
	for _, result := range results {
		if characterConfig.IsVerbose {
			fmt.Printf("%30s -> %d characters\n", result.FilePath, result.Count)
		}
		total += result.Count
	}
	fmt.Printf("\n%s has %d total characters\n", file, total)
	return err
}

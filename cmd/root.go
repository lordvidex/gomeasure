package cmd

/*
Copyright © 2022 Evans Owamoyo <evans.dev99@gmail.com>

*/

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/lordvidex/gomeasure/pkg/gomeasure"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var (
	// cfgFile is a manually overridden generalConfig file provided by the user
	cfgFile string

	// flags should be parsed into generalConfig struct
	generalConfig = gomeasure.NewConfig()

	rootCmd = &cobra.Command{
		Use:   "gomeasure",
		Short: "gomeasure is a CLI tool that provides quantitative analysis of a project",
		Long: `gomeasure is a CLI tool that provides quantitative analysis of a project.

it can be used to count the number of files in a directory recursively parsing subdirectories,
count the number of lines in all files of a directory, 
and count the number of characters in all files of a directory.

It also includes various flags that can be used to customize the output of the tool.
Run 'gomeasure --help' to see the available flags.`,
		Version: "0.3.1",
	}
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
	cobra.OnInitialize(initRootConfig)

	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "generalConfig file (default is $HOME/.gomeasure.yaml)")

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "displays the files being processed and their line count when `true`, default value `false`")
	rootCmd.PersistentFlags().StringP("include", "i", "", "include files that matches a given glob pattern e.g. `*.go`, `**/*.py`")
	rootCmd.PersistentFlags().StringP("no-include", "I", "", "exclude files that matches a given glob pattern e.g. `.git/**`, `.gitignore` or lists of files e.g. '{.git/**,.gitignore}' WITHOUT spaces")
}

// initRootConfig reads in generalConfig file searching the working directory and the home directory
func initRootConfig() {
	if cfgFile != "" {
		// Use generalConfig file from the flag.
		viper.SetConfigFile(cfgFile)
		cobra.CheckErr(viper.ReadInConfig())
	} else {
		viper.SetConfigName(".gomeasure")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")

		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search generalConfig in home directory with name ".gomeasure" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gomeasure")
		_ = viper.ReadInConfig()
	}

	// If a generalConfig file is found, read it in.
	data := viper.GetStringMap("general")
	bytes, err := json.Marshal(data)
	cobra.CheckErr(err)
	cobra.CheckErr(json.Unmarshal(bytes, &generalConfig))
}

func parseFlags(cmd *cobra.Command, config *gomeasure.Config) {
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if !flag.Changed {
			return
		}
		switch flag.Name {
		case "verbose":
			config.IsVerbose = flag.Value.String() == "true"
		case "include":
			config.IncludedFiles = flag.Value.String()
		case "no-include":
			config.ExcludedFiles = flag.Value.String()
		case "empty":
			config.ShouldCountEmpty = flag.Value.String() == "true"
		case "workers":
			intVal, err := strconv.Atoi(flag.Value.String())
			if err != nil {
				return
			}
			if intVal > 0 {
				config.WorkersCount = intVal
			}
		}
	})
}

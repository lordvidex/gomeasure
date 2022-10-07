package cmd

import (
	"github.com/lordvidex/gomeasure/pkg"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [version_number]",
	Short: "update gomeasure cli to another version(upgrade or downgrade)",

	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(update())
	},
}

func update() error {
	runner := &pkg.Runner{}
	err := runner.Update(rootCmd.Version)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

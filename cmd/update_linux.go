package cmd

import (
	"github.com/lordvidex/gomeasure/pkg"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update ",
	Short: "update gomeasure cli to latest version",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(update())
	},
}

func update() error {
	updater := &pkg.Updater{}
	err := updater.Update(rootCmd.Version)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

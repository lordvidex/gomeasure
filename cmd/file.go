package cmd

/*
Copyright © 2022 Evans Owamoyo <evans.dev99@gmail.com>
*/
import (
	"fmt"
	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use: "file",

	Short: "processes the number of files in a directory",
	Long:  `gomeasure file processes and returns the number of files in a directory / project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(len(cmd.PersistentFlags().Args()), len(cmd.Flags().Args()))
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

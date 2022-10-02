package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

type Message struct {
	Name string `json:"name"`
	//ZipballUrl string `json:"zipball_url"`
	//TarballUrl string `json:"tarball_url"`
	//Commit     struct {
	//	Sha string `json:"sha"`
	//	Url string `json:"url"`
	//} `json:"commit"`
	//NodeId string `json:"node_id"`
}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [version_number]",
	Short: "update gomeasure cli to another version(upgrade or downgrade)",

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			cobra.CheckErr(errors.New("passed to many arguments"))
		} else if len(args) == 0 { // no explicit version is passed
			cobra.CheckErr(update(""))
		} else if args[0] == "show" { // command to show all versions
			cobra.CheckErr(checkUpdate())
		} else { // update to latest version if exists
			cobra.CheckErr(update(args[0]))
		}
	},
}

func update(explicitVersion string) error {
	return errors.New("new thing error")
}

func checkUpdate() error {
	versions, err := getUpdateVersions()
	if err != nil {
		return err
	}
	fmt.Println("Current Available versions:")
	for _, element := range versions {
		fmt.Println(element)
	}
	return nil
}

func getUpdateVersions() ([]string, error) {
	url := "https://api.github.com/repos/lordvidex/gomeasure/tags"
	resp, err := http.Get(url)
	if err != nil {
		return []string{}, fmt.Errorf("cannot fetch URL %q: %v", url, err)
	}
	defer resp.Body.Close()

	var messages []Message

	err = json.NewDecoder(resp.Body).Decode(&messages)
	names := make([]string, len(messages))
	for i := 0; i < len(messages); i++ {
		names[i] = messages[i].Name
	}

	return names, nil
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

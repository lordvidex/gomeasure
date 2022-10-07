package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/exec"
	"strings"
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

	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(update())
	},
}

func update() error {
	latest, err := getVersions()
	if err != nil {
		return err
	}
	if "v"+rootCmd.Version == latest {
		fmt.Println("gomeasure is up to date")
	} else {
		fmt.Printf("Version %s is avaliable, install new verison?(y/n)", latest)
		reader := bufio.NewReader(os.Stdin)
		for true {
			input, _ := reader.ReadString('\n')
			input = strings.ToLower(input[:len(input)-1])
			if input == "y" {
				err = updateCli()
				if err != nil {
					return err
				}
				break
			} else if input == "n" {
				fmt.Println("Aborted update process")
				break
			} else {
				fmt.Println("Unknown input please try again")
			}
		}
	}
	return nil
}

func updateCli() error {
	branch := "apt-deploy"
	link := "curl \"https://raw.githubusercontent.com/lordvidex/gomeasure/" + branch + "/scripts/install.sh\" | sh"
	cmd := exec.Command("bash", "-c", link)
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

func getVersions() (string, error) {
	url := "https://api.github.com/repos/lordvidex/gomeasure/tags"
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("cannot fetch URL %q: %v", url, err)
	}
	defer resp.Body.Close()

	var messages []Message

	err = json.NewDecoder(resp.Body).Decode(&messages)
	latest := messages[len(messages)-1].Name

	return latest, nil
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

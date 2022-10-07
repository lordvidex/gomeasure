package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func (r *Runner) Update(currentVersion string) error {
	latest, err := r.getVersions()
	if err != nil {
		return err
	}

	if "v"+currentVersion == latest {
		fmt.Println("gomeasure is up to date")
	} else {
		fmt.Printf("Version %s is avaliable, install new verison?(y/n)", latest)
		reader := bufio.NewReader(os.Stdin)
		for true {
			input, _ := reader.ReadString('\n')
			input = strings.ToLower(input[:len(input)-1])
			if input == "y" {
				err = r.updateCli()
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

func (r *Runner) updateCli() error {
	branch := "master"
	link := "curl \"https://raw.githubusercontent.com/lordvidex/gomeasure/" + branch + "/scripts/install.sh\" | sh"
	cmd := exec.Command("bash", "-c", link)
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

func (r *Runner) getVersions() (string, error) {
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

package pkg

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Updater is the struct that handles the update process
// for convenience caches should be made to prevent checking versions
// too regularly
type Updater struct{}

func (u *Updater) Update(currentVersion string) error {
	latest, err := u.getLatestVersion()
	if err != nil {
		return err
	}
	if "v"+currentVersion == latest {
		fmt.Println("gomeasure is up to date")
		result, err := exec.Command("gomeasure", "--version").Output()
		if err != nil {
			return err
		}
		fmt.Println(string(result))
	} else {
		fmt.Printf("Version %s is avaliable, install new verison?(y/n)", latest)
		reader := bufio.NewReader(os.Stdin)

		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(strings.ToLower(input))
			if input == "y" {
				err = u.updateCli()
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

func (u *Updater) updateCli() error {
	cmd := exec.Command("bash", "-c", "curl \"https://raw.githubusercontent.com/lordvidex/gomeasure/master/scripts/install.sh\" | sh")
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

func (u *Updater) getLatestVersion() (s string, err error) {
	cmd := exec.Command("bash", "-c", "git ls-remote --refs --sort=\"version:refname\" --tags \"https://github.com/lordvidex/gomeasure\"  | cut -d/ -f3-|tail -n1\n")
	version, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(version)), nil
}

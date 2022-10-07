package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gobwas/glob"
	"github.com/pkg/errors"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

// Result is the result of a single file count
// it contains the file path and the number of lines
type Result struct {
	FilePath   string
	FileName   string
	LinesCount int64
}

// Runner is the main struct that processes the files in a folder
type Runner struct {
	// list of all file paths found in the Directory provided
	result []*Result

	// regex expression of files to be included on the command count process
	IncludedFiles string

	// regex expression of files to be excluded on the command count process
	ExcludedFiles string

	// bool value to weather to count empty lines in the command count process
	ShouldCountEmpty bool

	// number of concurrent workers used to read files and count them
	WorkersCount int

	// the Directory to be processed
	Directory string

	// should count the number of lines in the files
	ShouldCountLines bool
}

// Run reads the files, coordinates the workers and returns the result array / errors if any
// this function is safe i.e. it doesn't throw when an error occurs so that the program can
// report all errors at once to the user.
func (r *Runner) Run() ([]*Result, error) {
	err := r.readFiles()
	if err != nil {
		return nil, err
	}
	if !r.ShouldCountLines {
		return r.result, nil
	}

	// divide the files (jobs) into groups of workersCount
	nt := len(r.result) / r.WorkersCount
	wg := &sync.WaitGroup{}
	wg.Add(r.WorkersCount)
	errs := make([]error, 0)
	for i := 0; i < r.WorkersCount-1; i++ {
		i := i
		go func() {
			err := r.process(r.result[i*nt:(i+1)*nt], wg)
			if err != nil {
				errs = append(errs, err)
			}
		}()
	}
	go func() {
		err := r.process(r.result[(r.WorkersCount-1)*nt:], wg)
		if err != nil {
			errs = append(errs, err)
		}
	}()
	wg.Wait()
	return r.result, concatErrors(errs)
}

// countLines opens the file and counts the number of LinesCount in a file
func (r *Runner) countLines(path string, name string) (int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var count int64

	for scanner.Scan() {
		if !r.ShouldCountEmpty &&
			len(strings.Trim(scanner.Text(), "\n\r")) == 0 {
			continue
		}
		count++
	}

	return count, nil
}

// ReadFiles reads all files in the current Directory and all subdirectories
// and stores them in filePaths
func (r *Runner) readFiles() error {
	include, err := glob.Compile(r.IncludedFiles)
	exclude, err := glob.Compile(r.ExcludedFiles)
	if err != nil {
		return err
	}
	return filepath.WalkDir(r.Directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			isIncluded := len(r.IncludedFiles) == 0 || include.Match(path) || include.Match(d.Name())
			isExcluded := len(r.ExcludedFiles) != 0 && (exclude.Match(path) || exclude.Match(d.Name()))
			if isIncluded && !isExcluded {
				r.result = append(r.result, &Result{FilePath: path, FileName: d.Name()})
			}
		}
		return nil
	})
}

// process is a function meant to be run in a goroutine
// it receives arr []*Result which contains the files to be processed
// and optional wg *sync.WaitGroup which sends a Done signal when the function is done
func (r *Runner) process(arr []*Result, wg *sync.WaitGroup) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()
	for _, result := range arr {
		lines, err := r.countLines(result.FilePath, result.FileName)
		if err != nil {
			return err
		}
		result.LinesCount = lines
	}
	return nil
}

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
	branch := "apt-deploy"
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

func concatErrors(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	err := errs[0]
	for _, e := range errs[1:] {
		err = errors.Wrap(err, e.Error())
	}
	return err
}

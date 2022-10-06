package gomeasure

import (
	"bufio"
	"github.com/gobwas/glob"
	"github.com/pkg/errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// measureAction determines the type of measurement to be performed
type measureAction int

const (
	// MeasureFile counts the number of files in a project or directory
	MeasureFile measureAction = iota

	// MeasureLine counts the number of lines in a file
	MeasureLine

	// MeasureCharacter counts the number of characters in a file (files inside a folder)
	MeasureCharacter
)

// Result is the result of a single file count
// it contains the file path and the number of lines
type Result struct {
	FilePath string
	FileName string
	Count    int64
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

	// specifies whether to count files, lines or characters
	Action measureAction
}

// Run reads the files, coordinates the workers and returns the result array / errors if any
// this function is safe i.e. it doesn't throw when an error occurs so that the program can
// report all errors at once to the user.
func (r *Runner) Run() ([]*Result, error) {
	err := r.readFiles()
	if err != nil {
		return nil, err
	}
	if r.Action == MeasureFile {
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

// countLines opens the file and counts the number of Count in a file
func (r *Runner) countLines(path string, name string) (int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var count int64

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if r.ShouldCountEmpty ||
			len(text) != 0 {
			if r.Action == MeasureCharacter {
				count += int64(len([]byte(text)))
			} else if r.Action == MeasureLine {
				count++
			}
		}
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
		result.Count = lines
	}
	return nil
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

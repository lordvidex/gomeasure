package pkg

import (
	"bufio"
	"github.com/pkg/errors"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

// Result is the result of a single file count
// it contains the file path and the number of lines
type Result struct {
	file  string
	lines int64
}

// Runner is the main struct that processes the files in a folder
type Runner struct {
	// list of all file paths found in the directory provided
	result []*Result
	// number of concurrent workers used to read files and count them
	workersCount int
	// the directory to be processed
	directory string
}

func Default() *Runner {
	return &Runner{
		workersCount: 5,
		directory:    ".",
	}
}

func New(directory string, workers int) *Runner {
	return &Runner{
		directory:    directory,
		workersCount: workers,
	}
}

// Run reads the files, coordinates the workers and returns the result array / errors if any
// this function is safe i.e. it doesn't throw when an error occurs so that the program can
// report all errors at once to the user.
func (r *Runner) Run() ([]*Result, error) {
	err := r.readFiles()
	if err != nil {
		return nil, err
	}
	// divide the files (jobs) into groups of workersCount
	nt := len(r.result) / r.workersCount
	wg := &sync.WaitGroup{}
	errs := make([]error, 0)
	for i := 0; i < r.workersCount-1; i++ {
		i := i
		go func() {
			err := r.process(r.result[i*nt:(i+1)*nt], wg)
			if err != nil {
				errs = append(errs, err)
			}
		}()
	}
	go func() {
		err := r.process(r.result[(r.workersCount-1)*nt:], wg)
		if err != nil {
			errs = append(errs, err)
		}
	}()
	return r.result, concatErrors(errs)
}

// countLines opens the file and counts the number of lines in a file
func countLines(path string) (int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	var count int64
	for scanner.Scan() {
		count++
	}

	return count, nil
}

// ReadFiles reads all files in the current directory and all subdirectories
// and stores them in filePaths
func (r *Runner) readFiles() error {
	return filepath.WalkDir(r.directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			// TODO:FEATURE - add filters for file extensions in Runner
			r.result = append(r.result, &Result{file: path})
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
		lines, err := countLines(result.file)
		if err != nil {
			return err
		}
		result.lines = lines
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

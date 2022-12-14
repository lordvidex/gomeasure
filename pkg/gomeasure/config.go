package gomeasure

import (
	"encoding/json"
	"strings"
)

const (
	defaultWorkersCount = 5
)

// Config is the configuration struct for gomeasure.
type Config struct {
	// glob expression of files to be included on the command count process
	IncludedFiles string `json:"include"`

	// glob expression of files to be excluded on the command count process
	ExcludedFiles string `json:"exclude"`

	// bool value to weather to count empty lines in the command count process
	ShouldCountEmpty bool `json:"empty_lines"`

	// number of concurrent workers used to read files and count them
	WorkersCount int `json:"workers"`

	// IsVerbose determines if the program should print the progress of the command count process
	IsVerbose bool `json:"verbose"`
}

// NewConfig returns a new Config with default values
func NewConfig() *Config {
	return &Config{
		WorkersCount: defaultWorkersCount,
	}
}

func (c *Config) UnmarshalJSON(data []byte) error {
	type temp struct {
		Include []string
		Exclude []string
		Empty   bool
		Workers int
		Verbose bool
	}
	t := temp{}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	if len(t.Exclude) != 0 {
		c.ExcludedFiles = sliceToGlobString(t.Exclude)
	}
	if len(t.Include) != 0 {
		c.IncludedFiles = sliceToGlobString(t.Include)
	}
	if t.Verbose {
		c.IsVerbose = t.Verbose
	}
	if t.Empty {
		c.ShouldCountEmpty = t.Empty
	}
	if t.Workers > 0 {
		c.WorkersCount = t.Workers
	}
	return nil
}

// sliceToGlobString converts a slice of strings to a glob string
// e.g. []string{"*.go", "*.txt"} -> "{*.go,*.txt}"
func sliceToGlobString(arr []string) string {
	if len(arr) == 0 {
		return ""
	} else if len(arr) == 1 {
		return arr[0]
	} else {
		// trim each element and join them with a comma
		arr = func(input []string) []string {
			output := make([]string, 0, len(input))
			for _, v := range input {
				trimmed := strings.TrimSpace(v)
				if len(trimmed) != 0 {
					output = append(output, trimmed)
				}
			}
			return output
		}(arr)
		return "{" + strings.Join(arr, ",") + "}"
	}
}

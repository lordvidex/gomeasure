package cmd

import (
	"github.com/lordvidex/gomeasure/pkg/gomeasure"
	"testing"
)

var (
	testDir = "../testdata"
)

func computeResultsTotal(result []*gomeasure.Result) int64 {
	var total int64 = 0
	for _, r := range result {
		total += r.Count
	}
	return total
}

func Test_processCharacters(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		wantTotal int64
	}{
		{
			"test character count with lines having whitespaces",
			args{testDir + "/whitespace.golden"},
			false,
			4,
		},
		{
			"test character count for files within nested folders",
			args{testDir + "/testfolder"},
			false,
			37 + 74,
		},
		{
			"test folder not found",
			args{testDir + "folder_y"},
			true,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := &gomeasure.Runner{
				IncludedFiles:    include,
				ExcludedFiles:    exclude,
				ShouldCountEmpty: countEmptyLines,
				WorkersCount:     workersCount,
				Directory:        tt.args.file,
				Action:           gomeasure.MeasureCharacter,
			}
			results, err := runner.Run()
			if (err != nil) != tt.wantErr {
				t.Errorf("processCharacters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testTotal := computeResultsTotal(results)
			if testTotal != tt.wantTotal {
				t.Errorf("processCharacters() total = %v, want %v", testTotal, tt.wantTotal)
			}

		})
	}
}

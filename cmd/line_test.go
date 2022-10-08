package cmd

import (
	"github.com/lordvidex/gomeasure/pkg/gomeasure"
	"testing"
)

func Test_processLines(t *testing.T) {
	type args struct {
		directory       string
		include         string
		exclude         string
		countEmptyLines bool
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		wantTotal int64
		wantFiles int
	}{
		{
			"test include golden file and count empty lines",
			args{
				testDir,
				"*.golden",
				"",
				true,
			},
			false,
			13,
			2,
		},
		{
			"test don't count empty lines",
			args{
				testDir,
				"",
				"",
				false,
			},
			false,
			13,
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := &gomeasure.Runner{
				Config: &gomeasure.Config{
					IncludedFiles:    tt.args.include,
					ExcludedFiles:    tt.args.exclude,
					ShouldCountEmpty: tt.args.countEmptyLines,
					WorkersCount:     1,
				},
				Directory: tt.args.directory,
				Action:    gomeasure.MeasureLine,
			}
			results, err := runner.Run()

			if (err != nil) != tt.wantErr {
				t.Errorf("processLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testTotal := computeResultsTotal(results)
			if testTotal != tt.wantTotal {
				t.Errorf("number of total lines expected = %v, got %v", tt.wantTotal, testTotal)
				return
			}
			if len(results) != tt.wantFiles {
				t.Errorf("number of files expected = %v, got %v", tt.wantFiles, len(results))
			}
		})
	}
}

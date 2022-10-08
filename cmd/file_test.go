package cmd

import (
	"github.com/lordvidex/gomeasure/pkg/gomeasure"
	"testing"
)

func Test_processFiles(t *testing.T) {
	type args struct {
		directory string
		include   string
		exclude   string
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		filesCount int
	}{
		{
			"test include golden file",
			args{"../testdata", "*.golden", ""},
			false,
			2,
		},
		{
			"default",
			args{"../testdata", "", ""},
			false,
			3,
		},
		{
			"exclude golden files",
			args{"../testdata", "", "*.golden"},
			false,
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := &gomeasure.Runner{
				Config: &gomeasure.Config{
					IncludedFiles: tt.args.include,
					ExcludedFiles: tt.args.exclude,
				},
				Directory: tt.args.directory,
				Action:    gomeasure.MeasureFile,
			}
			results, err := runner.Run()
			if (err != nil) != tt.wantErr {
				t.Errorf("processFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(results) != tt.filesCount {
				t.Errorf("processFiles() expected = %v, got %v", tt.filesCount, len(results))
			}
		})
	}
}

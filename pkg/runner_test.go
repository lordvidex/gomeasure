package pkg

import (
	"github.com/pkg/errors"
	"strings"
	"sync"
	"testing"
)

func TestRunner_process(t *testing.T) {
	type fields struct {
		result       []*Result
		workersCount int
		directory    string
	}
	type args struct {
		arr []*Result
		wg  *sync.WaitGroup
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"nil waitgroup",
			fields{},
			args{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Runner{
				result:       tt.fields.result,
				WorkersCount: tt.fields.workersCount,
				Directory:    tt.fields.directory,
			}
			if err := r.process(tt.args.arr, tt.args.wg); (err != nil) != tt.wantErr {
				t.Errorf("process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_concatErrors(t *testing.T) {
	type args struct {
		errs []error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"empty error struct",
			args{make([]error, 0)},
			false,
		},
		{
			"two errors",
			args{[]error{
				errors.New("error 2"),
				errors.New("error 1"),
			}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := concatErrors(tt.args.errs)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("there shouldn't be an error")
				} else {
					for _, e := range tt.args.errs {
						if !strings.Contains(err.Error(), e.Error()) {
							t.Errorf("err does not wrap expected error %s", e.Error())
						}
					}
				}
			} else {
				if tt.wantErr {
					t.Errorf("Error is expected but function returned nil")
				}
			}
		})
	}
}

package gomeasure

import "testing"

func TestSliceToGlobString(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{[]string{}}, ""},
		{"one", args{[]string{"*.go"}}, "*.go"},
		{"multiple", args{[]string{"*.go", "*.txt"}}, "{*.go,*.txt}"},
		{"multiple with empty", args{[]string{"*.go", "*.txt", ""}}, "{*.go,*.txt}"},
		{"multiple with spaces", args{[]string{" *.go", "*.txt", " "}}, "{*.go,*.txt}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceToGlobString(tt.args.arr); got != tt.want {
				t.Errorf("sliceToGlobString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConfig(t *testing.T) {
	c := NewConfig()
	if c.IsVerbose {
		t.Errorf("NewConfig() IsVerbose = %v, want %v", c.IsVerbose, false)
	}
	if c.ShouldCountEmpty {
		t.Errorf("NewConfig() ShouldCountEmpty = %v, want %v", c.ShouldCountEmpty, false)
	}
	if c.WorkersCount != defaultWorkersCount {
		t.Errorf("NewConfig WorkersCount = %v, want %v", c.WorkersCount, defaultWorkersCount)
	}
}

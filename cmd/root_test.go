package cmd

import (
	"fmt"
	"testing"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Execute()
		})
	}
}

func Test_initConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"yay",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//cfgFile = "../example/.gomeasure.yaml"
			initConfig()
			fmt.Printf("%+v", generalConfig)
		})
	}
}

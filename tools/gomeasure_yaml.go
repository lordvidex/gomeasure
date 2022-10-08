package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/invopop/jsonschema"
)

type Flags struct {
	Verbose bool     `json:"verbose,omitempty" jsonschema:"description=verbose output,default=false"`
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

type LineFlags struct {
	Flags
	EmptyLines bool `json:"empty_lines,omitempty" jsonschema:"description=include empty lines,default=false"`
	Workers    int  `json:"workers,omitempty" jsonschema:"description=number of workers,default=5"`
}

type GOMeasureYAML struct {
	General   Flags     `json:"general,omitempty"`
	Line      LineFlags `json:"line,omitempty"`
	File      Flags     `json:"file,omitempty"`
	Character LineFlags `json:"character,omitempty"`
}

func main() {
	s := jsonschema.Reflect(GOMeasureYAML{})

	data, err := json.MarshalIndent(s, "", "  ")
	file, err := os.Create("docs/gomeasure.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, string(data))
}

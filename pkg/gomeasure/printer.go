package gomeasure

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
)

// PrettyPrint prints rows of lines in a tabular format.
// Each line is a row
// and each column of line is separated by a tab.
func PrettyPrint(lines []string) string {
	var buf bytes.Buffer
	writer := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	for _, line := range lines {
		if !strings.HasSuffix(line, "\t") {
			line += "\t"
		}
		fmt.Fprintln(writer, line)
	}
	writer.Flush()
	return buf.String()
}

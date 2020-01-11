package goal_test

import (
	"bytes"
	"fmt"

	"github.com/sweetycode/goal"
)

func ExampleScanLines() {
	s := `
# goal Document

goal is libaray for daily golang development. Hope provide taste of Python.
	`

	for line := range goal.ScanLines(bytes.NewBufferString(s)) {
		fmt.Println(line)
	}
	//Output:
	//# goal Document
	//goal is libaray for daily golang development. Hope provide taste of Python.
}

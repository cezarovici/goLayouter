package item

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestKindOfFile is a unit test function to test the KindOfFile function
// It tests the KindOfFile function with different inputs and expected
// outputs using a table-driven approach.
func TestKindOfFile(t *testing.T) {
	type testCase struct {
		test   string
		input  string
		output KindOfFile
	}

	testCases := []testCase{
		{
			test:   "test main",
			input:  "main.go",
			output: Main,
		},
		{
			test:   "test file",
			input:  "func_test.go",
			output: Test,
		},
		{
			test:   "object file",
			input:  "obj_file.go",
			output: Object,
		},
		{
			test:   "object test",
			input:  "obj_file_test.go",
			output: Test,
		},
		{
			test:   "normal file",
			input:  "file.go",
			output: NormalFile,
		},
		{
			test:   "input folder",
			input:  "folder1",
			output: Folder,
		},
	}

	for _, currentTestCase := range testCases {
		currentTestCase := currentTestCase

		t.Run(currentTestCase.test, func(t *testing.T) {
			require.Equal(t, currentTestCase.output,
				NewKindOfFile(currentTestCase.input))
		})
	}
}

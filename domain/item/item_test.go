package item_test

import (
	"testing"

	"github.com/cezarovici/goLayouter/domain/item"
	"github.com/stretchr/testify/require"
)

// TestKindOfFile is a unit test function to test the KindOfFile function
// It tests the KindOfFile function with different inputs and expected
// outputs using a table-driven approach.
func TestKindOfFile(t *testing.T) {
	t.Parallel()

	type testCase struct {
		test   string
		input  string
		output item.KindOfFile
	}

	testCases := []testCase{
		{
			test:   "test main",
			input:  "main.go",
			output: item.Main,
		},
		{
			test:   "test file",
			input:  "func_test.go",
			output: item.Test,
		},
		{
			test:   "object file",
			input:  "obj_file.go",
			output: item.Object,
		},
		{
			test:   "object test",
			input:  "obj_file_test.go",
			output: item.Test,
		},
		{
			test:   "normal file",
			input:  "file.go",
			output: item.NormalFile,
		},
		{
			test:   "input folder",
			input:  "folder1",
			output: item.Folder,
		},
	}

	for _, currentTestCase := range testCases {
		currentTestCase := currentTestCase

		t.Run(currentTestCase.test, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, currentTestCase.output,
				item.NewKindOfFile(currentTestCase.input))
		})
	}
}

package lineparser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetObjectName(t *testing.T) {
	var line lineParser

	type testCase struct {
		test   string
		input  string
		output lineParser
	}

	testCases := []testCase{
		{
			test:  "simple object file",
			input: "app/test/obj_file.go",
			output: lineParser{
				ObjectName: "File",
			},
		},
		{
			test:   "simple test file",
			input:  "app/test/file_test.go",
			output: lineParser{},
		},
		{
			test:  "obj test file",
			input: "app/test/obj_file_test.go",
			output: lineParser{
				ObjectName: "File",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			line.SetObjectName(tc.input)
			require.Equal(t, tc.output, line)
		})
	}
}

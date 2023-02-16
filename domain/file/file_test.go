package file

import (
	"os"
	"testing"

	"github.com/cezarovici/goLayouter/app/helpers"
	"github.com/stretchr/testify/require"
)

func TestWriteToDisk(t *testing.T) {
	type testCase struct {
		test    string
		input   File
		content bool

		errorExpected error
	}

	testCases := []testCase{
		{
			test: "file without content",
			input: File{
				path: "test.go",
			},
			content:       false,
			errorExpected: nil,
		},
		{
			test: "file with content",
			input: File{
				path:    "main.go",
				content: "#package main",
			},
			content: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.errorExpected, tc.input.WriteToDisk())

			if tc.content {
				outputContent, errRead := helpers.ReadFile(tc.input.path)
				require.Equal(t, outputContent[0], tc.input.content)
				require.NoError(t, errRead)
			}

			require.NoError(t, os.Remove(tc.input.path))
		})
	}
}

package file

import (
	"io/ioutil"
	"os"
	"testing"

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
				Path: "test.go",
			},
			content:       false,
			errorExpected: nil,
		},
		{
			test: "file with content",
			input: File{
				Path:    "main.go",
				Content: "#package main",
			},
			content:       true,
			errorExpected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			n, err := tc.input.Write(nil)
			require.Equal(t, tc.errorExpected, err)
			require.Equal(t, len(tc.input.Content), n)

			_, errStat := os.Stat(tc.input.Path)
			require.NoError(t, errStat)

			outputContent, errRead := ioutil.ReadFile(tc.input.Path)
			require.NoError(t, errRead)
			require.Equal(t, tc.input.Content, string(outputContent))
			require.NoError(t, os.Remove(tc.input.Path))

		})
	}
}

package file

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestWrite is a unit test that verifies the behavior of the Write() method
// of the File struct when writing content to disk.
func TestWrite(t *testing.T) {
	// Define a test case struct that contains the necessary information to run the test.
	type testCase struct {
		test          string // Name of the test case.
		input         File   // Input data for the test.
		content       bool   // Indicates whether the input file contains content or not.
		errorExpected error  // Expected error returned by the Write() method.
	}

	// Define the test cases to run.
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

	// Iterate over each test case and run the test.
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			// Call the Write() method and verify its return values.
			n, err := tc.input.Write(nil)
			require.Equal(t, tc.errorExpected, err)
			require.Equal(t, len(tc.input.Content), n)

			// Verify that the file was created and contains the expected content.
			_, errStat := os.Stat(tc.input.Path)
			require.NoError(t, errStat)

			outputContent, errRead := ioutil.ReadFile(tc.input.Path)
			require.NoError(t, errRead)
			require.Equal(t, tc.input.Content, string(outputContent))

			// Clean up by deleting the test file.
			require.NoError(t, os.Remove(tc.input.Path))
		})
	}
}

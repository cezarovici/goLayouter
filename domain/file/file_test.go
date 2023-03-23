package file_test

import (
	"os"
	"testing"

	"github.com/cezarovici/goLayouter/domain/file"
	"github.com/stretchr/testify/require"
)

// TestWrite is a unit test that verifies the behavior of the Write() method
// of the file.File struct when writing Package to disk.
func TestWrite(t *testing.T) {
	t.Parallel()
	// Define a test case struct that contains the necessary information to run the test.
	type testCase struct {
		test          string    // Name of the test case.
		input         file.File // Input data for the test.
		Package       bool      // Indicates whether the input file.file contains Package or not.
		errorExpected error     // Expected error returned by the Write() method.
	}

	// Define the test cases to run.
	testCases := []testCase{
		{
			test: "file.file without Package",
			input: file.File{
				Path: "test.go",
			},
			Package:       false,
			errorExpected: nil,
		},
		{
			test: "file.file with Package",
			input: file.File{
				Path:    "main.go",
				Package: "#package main",
			},
			Package:       true,
			errorExpected: nil,
		},
	}

	// Iterate over each test case and run the test.
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			// Call the Write() method and verify its return values.
			err := tc.input.Write(nil)
			require.Equal(t, tc.errorExpected, err)

			// Verify that the file.file was created and contains the expected Package.
			_, errStat := os.Stat(tc.input.Path)
			require.NoError(t, errStat)

			outputPackage, errRead := os.ReadFile(tc.input.Path)
			require.NoError(t, errRead)
			require.Equal(t, tc.input.Package, string(outputPackage))

			// Clean up by deleting the test file.file.
			require.NoError(t, os.Remove(tc.input.Path))
		})
	}
}

func TestGetContent(t *testing.T) {
	t.Parallel()

	type testCase struct {
		test  string
		input file.File

		output []byte
	}

	testCases := []testCase{
		{
			test:   "just 1 file.file",
			input:  file.File{Package: "main.go"},
			output: []byte("main.go"),
		},
		{
			test:   "file.file from previous path",
			input:  file.File{Package: "../main.go"},
			output: []byte("../main.go"),
		},
		{
			test:   "file.file from next path",
			input:  file.File{Package: "file.file/main.go"},
			output: []byte("file.file/main.go"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.output, tc.input.GetContent())
		})
	}
}

func TestGetPath(t *testing.T) {
	t.Parallel()

	type testCase struct {
		test  string
		input file.File

		output string
	}

	testCases := []testCase{
		{
			test:   "just 1 file.file",
			input:  file.File{Path: "main.go"},
			output: "main.go",
		},
		{
			test:   "file.file from previous path",
			input:  file.File{Path: "../main.go"},
			output: "../main.go",
		},
		{
			test:   "file.file from next path",
			input:  file.File{Path: "file.file/main.go"},
			output: "file.file/main.go",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.output, tc.input.GetPath())
		})
	}
}

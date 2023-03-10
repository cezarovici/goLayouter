package file

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestWrite is a unit test that verifies the behavior of the Write() method
// of the File struct when writing Package to disk.
func TestWrite(t *testing.T) {
	// Define a test case struct that contains the necessary information to run the test.
	type testCase struct {
		test          string // Name of the test case.
		input         File   // Input data for the test.
		Package       bool   // Indicates whether the input file contains Package or not.
		errorExpected error  // Expected error returned by the Write() method.
	}

	// Define the test cases to run.
	testCases := []testCase{
		{
			test: "file without Package",
			input: File{
				Path: "test.go",
			},
			Package:       false,
			errorExpected: nil,
		},
		{
			test: "file with Package",
			input: File{
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
			n, err := tc.input.Write(nil)
			require.Equal(t, tc.errorExpected, err)
			require.Equal(t, len(tc.input.Package), n)

			// Verify that the file was created and contains the expected Package.
			_, errStat := os.Stat(tc.input.Path)
			require.NoError(t, errStat)

			outputPackage, errRead := ioutil.ReadFile(tc.input.Path)
			require.NoError(t, errRead)
			require.Equal(t, tc.input.Package, string(outputPackage))

			// Clean up by deleting the test file.
			require.NoError(t, os.Remove(tc.input.Path))
		})
	}
}

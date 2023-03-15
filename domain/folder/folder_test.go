package folder

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestWrite is a unit test function for the Write() method of the Folder struct.
func TestWrite(t *testing.T) {
	// Define a struct to represent a test case
	type testCase struct {
		test          string // A description of the test case
		input         Folder // The input Folder object for the test
		alreadyExists bool   // Whether the input Folder object already exists
		expectedError error  // The expected error returned by the Write() method
	}

	// Define a list of test cases
	testCases := []testCase{
		{
			test:          "already existing directory",
			input:         Folder{"../folder"}, // A Folder object with an existing path
			alreadyExists: true,
			expectedError: nil,
		},
		{
			test:          "new directory",
			input:         Folder{"newFolder"}, // A Folder object with a new path
			alreadyExists: false,
			expectedError: nil,
		},
	}

	// Loop through each test case
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			// Call the Write() method with a nil logger and capture the returned error
			err := tc.input.Write(nil)

			// Assert that the error returned by Write() matches the expected error
			require.Equal(t, tc.expectedError, err)

			// If the input Folder object doesn't exist yet, assert that it was successfully removed
			if !tc.alreadyExists {
				require.NoError(t, os.Remove(tc.input.Path))
			}
		})
	}
}

func TestGetPackage(t *testing.T) {
	folder := Folder{
		Path: "folder1",
	}

	require.NoError(t, folder.Write(nil))
	require.Equal(t, []uint8([]byte(nil)), folder.GetPackage())
	require.NoError(t, os.Remove(folder.Path))
}

package folder_test

import (
	"os"
	"testing"

	"github.com/cezarovici/goLayouter/domain/folder"
	"github.com/stretchr/testify/require"
)

// TestWrite is a unit test function for the Write() method of the folder.Folder struct.
func TestWrite(t *testing.T) {
	// Define a struct to represent a test case
	type testCase struct {
		test          string        // A description of the test case
		input         folder.Folder // The input folder.Folder object for the test
		alreadyExists bool          // Whether the input folder.Folder object already exists
		expectedError error         // The expected error returned by the Write() method
	}

	// Define a list of test cases
	testCases := []testCase{
		{
			test:          "already existing directory",
			input:         folder.Folder{"../folder.folder"}, // A folder.Folder object with an existing path
			alreadyExists: true,
			expectedError: nil,
		},
		{
			test:          "new directory",
			input:         folder.Folder{"newfolder.Folder"}, // A folder.Folder object with a new path
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

			// If the input folder.Folder object doesn't exist yet, assert that it was successfully removed
			if !tc.alreadyExists {
				require.NoError(t, os.Remove(tc.input.Path))
			}
		})
	}
}

func TestGetContent(t *testing.T) {
	folder := folder.Folder{
		Path: "folder.folder1",
	}

	require.NoError(t, folder.Write(nil))
	require.Equal(t, []byte(nil), folder.GetContent())
	require.NoError(t, os.Remove(folder.Path))
}

func TestGetPath(t *testing.T) {
	folder := folder.Folder{
		Path: "folder.folder1",
	}

	require.NoError(t, folder.Write(nil))
	require.Equal(t, "folder.folder1", folder.GetPath())
	require.NoError(t, os.Remove(folder.Path))
}

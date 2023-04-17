package folder_test

import (
	"os"
	"testing"

	"github.com/cezarovici/goLayouter/domain/folder"
	"github.com/stretchr/testify/require"
)

// TestWrite is a unit test function for the Write() method of the folder.Folder struct.
func TestWrite(t *testing.T) {
	t.Parallel()

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
			input:         folder.Folder{"../folder"}, // A folder.Folder object with an existing path
			alreadyExists: true,
			expectedError: nil,
		},
		{
			test:          "new directory",
			input:         folder.Folder{"newfolder"}, // A folder.Folder object with a new path
			alreadyExists: false,
			expectedError: nil,
		},
	}

	// Loop through each test case
	for _, currentTestCase := range testCases {
		currentTestCase := currentTestCase

		t.Run(currentTestCase.test, func(t *testing.T) {
			t.Parallel()

			// Call the Write() method with a nil logger and capture the returned error
			err := currentTestCase.input.Write(nil)

			// Assert that the error returned by Write() m currentTestCasehes the expected error
			require.Equal(t, currentTestCase.expectedError, err)

			// If the input folder.Folder object doesn't exist yet, assert that it was successfully removed
			if !currentTestCase.alreadyExists {
				require.NoError(t, os.Remove(currentTestCase.input.Path))
			}
		})
	}
}

func TestGetContent(t *testing.T) {
	t.Parallel()

	folder := folder.Folder{
		Path: "folder1",
	}

	require.NoError(t, folder.Write(nil))
	require.Equal(t, []byte(nil), folder.GetContent())
	defer require.NoError(t, os.Remove(folder.Path))
}

func TestGetPath(t *testing.T) {
	t.Parallel()

	folder := folder.Folder{
		Path: "folder1",
	}

	require.NoError(t, folder.Write(nil))
	require.Equal(t, "folder1", folder.GetPath())
	defer require.NoError(t, os.Remove(folder.Path))
}

package folder

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestWrite is a unit test function for the Write() method of the folder.Folder struct.
func TestWrite(t *testing.T) {
	// Define a struct to represent a test case
	type testCase struct {
		test          string // A description of the test case
		input         Folder // The input folder.Folder object for the test
		alreadyExists bool   // Whether the input folder.Folder object already exists
		expectedError error  // The expected error returned by the Write() method
	}

	// Define a list of test cases
	testCases := []testCase{
		{
			test:          "already existing directory",
			input:         Folder{"../folder"}, // A folder.Folder object with an existing path
			alreadyExists: true,
			expectedError: nil,
		},
		{
			test:          "new directory",
			input:         Folder{"newfolder"}, // A folder.Folder object with a new path
			alreadyExists: false,
			expectedError: nil,
		},
	}

	// Loop through each test case
	for _, currentTestCase := range testCases {
		currentTestCase := currentTestCase

		t.Run(currentTestCase.test, func(t *testing.T) {
			// Call the Write() method with a nil logger and capture the returned error
			err := currentTestCase.input.Write(nil)

			// Assert that the returned error is the same as the expected error
			require.Equal(t, currentTestCase.expectedError, err)

			// If the test case is for a new directory, remove the directory
			if !currentTestCase.alreadyExists {
				require.NoError(t, os.Remove(currentTestCase.input.Path))
			}
		})
	}
}

func TestGetContent(t *testing.T) {
	folder := Folder{
		Path: "folder1",
	}

	require.NoError(t, folder.Write(nil))
	require.Equal(t, []byte(nil), folder.GetContent())

	defer require.NoError(t, os.Remove(folder.Path))
}

func TestGetPath(t *testing.T) {
	folder1 := Folder{
		Path: "folder1",
	}

	folder2 := Folder{
		Path: "folder1/f3",
	}

	require.NoError(t, folder1.Write(nil))
	require.NoError(t, folder2.Write(nil))

	require.Equal(t, "folder1", folder1.GetPath())
	require.Equal(t, "folder1/f3", folder2.GetPath())

	defer require.NoError(t, os.Remove(folder2.Path))
	defer require.NoError(t, os.Remove(folder1.Path))
}

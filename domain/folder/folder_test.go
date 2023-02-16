package folder

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteToDisk(t *testing.T) {
	type testCase struct {
		test          string
		input         Folder
		alreadyExists bool

		expectedError error
	}

	testCases := []testCase{
		{
			test:          "already existing directory",
			input:         Folder{"../folder"},
			alreadyExists: true,
			expectedError: nil,
		},
		{
			test:          "new directory",
			input:         Folder{"newFolder"},
			alreadyExists: false,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.expectedError, tc.input.WriteToDisk())

			if !tc.alreadyExists {
				require.NoError(t, os.Remove(tc.input.path))
			}
		})
	}
}

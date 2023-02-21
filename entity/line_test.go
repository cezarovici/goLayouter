package line

import (
	"errors"
	"log"
	"testing"

	"github.com/cezarovici/goLayouter/app/helpers"
	"github.com/stretchr/testify/require"
)

func TestConvertToLine(t *testing.T) {
	type testCase struct {
		test   string
		input  string
		output Line
	}

	testCases := []testCase{
		{
			test:  "first line",
			input: "folder1",
			output: Line{
				info:  "folder1",
				level: 0,
			},
		},
		{
			test:  "different level",
			input: "  subfolder",
			output: Line{
				info:  "subfolder",
				level: 2,
			},
		},
		{
			test:  "package",
			input: " # package",
			output: Line{
				info:  "# package",
				level: 1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, ConvertToLine(tc.input))
		})
	}
}

func TestNewLines(t *testing.T) {
	type testCase struct {
		test   string
		input  []string
		output Lines

		errorExpected error
	}

	testCases := []testCase{
		// Cases with errors
		{
			test:          "no content",
			input:         nil,
			output:        nil,
			errorExpected: errors.New("no content parsed"),
		},

		// Happy cases
		{
			test:  "2 lines",
			input: []string{"folder1", " subfolder1"},
			output: Lines{
				Line{
					info:  "folder1",
					level: 0,
				},
				Line{
					info:  "subfolder1",
					level: 1,
				},
			},
			errorExpected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			lines, errCreatingLines := NewLines(tc.input)

			require.Equal(t, tc.errorExpected, errCreatingLines)
			require.Equal(t, tc.output, lines)
		})
	}
}

const _parseTestCases = "../testCases/parseTest/"

func TestToItems(t *testing.T) {
	type testCase struct {
		test   string
		input  string
		output string
	}

	testCases := []testCase{
		{
			test:   "folders with indents",
			input:  _parseTestCases + "foldersWithIndents/input",
			output: _parseTestCases + "foldersWithIndents/output",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			inputContent, errorReading := helpers.ReadFile(tc.input)
			log.Print(tc.input)
			require.NoError(t, errorReading)

			lines, errNewLines := NewLines(inputContent)
			require.NoError(t, errNewLines)

			outputContent, errorReading := helpers.ReadFile(tc.output)
			require.NoError(t, errorReading)

			items := lines.ToItems()
			require.Equal(t, outputContent, items.ToStrings())
		})
	}

}

package renders

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Model represents the data that will be used to render the templates.
type model struct {
	ObjectName string // Name of the main object in the Go file.
	Package    string // Name of the package the Go file belongs to.
}

// exampleModel is an instance of the Model struct that can be used for testing.
var exampleModel = model{
	ObjectName: "Entry",
	Package:    "package entry",
}

// TestRenderFuncs is a unit test that verifies the output of each rendering function.
func TestRenderFuncs(t *testing.T) {
	require.NoError(t, os.Chdir("../../cmd"))
	// Define a test case struct that contains the necessary information to run the test.
	type testCase struct {
		test           string // Name of the test case.
		kind           string // Name of the rendering function to test.
		outputTestName string // Path to the file that contains the expected output.
	}
	// Define the test cases to run.
	testCases := []testCase{
		{
			test:           "main render",
			kind:           "main",
			outputTestName: _mainOutputPath,
		},
		{
			test:           "test render",
			kind:           "test",
			outputTestName: _testOutputPath,
		},
		{
			test:           "obj render",
			kind:           "object",
			outputTestName: _objectOutputPath,
		},
		{
			test:           "tdd render",
			kind:           "tableDriven",
			outputTestName: _tddOutputPath,
		},
	}

	// Iterate over each test case and run the test.
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			buffer := "buffer"
			_, errCreating := os.Create(buffer)
			require.NoError(t, errCreating)

			// Read the expected output from the file system.
			bytesContent, errRead := os.ReadFile(tc.outputTestName)
			require.NoError(t, errRead)

			require.NoError(t, RenderFuncs[tc.kind](buffer, exampleModel))
			bytesExpected, errRead := os.ReadFile(buffer)
			require.NoError(t, errRead)

			// Verify that the output matches the expected output.
			require.Equal(t, bytesContent, bytesExpected)

			//Clean the buffer
			require.NoError(t, os.Remove(buffer))
		})
	}
}

package render_test

import (
	"os"
	"testing"

	"github.com/cezarovici/goLayouter/app/service/render"
	"github.com/cezarovici/goLayouter/domain/item"
	"github.com/stretchr/testify/require"
)

// Model represents the data that will be used to render the templates.
type model struct {
	ObjectName string // Name of the main object in the Go file.
	Package    string // Name of the package the Go file belongs to.
}

// TestRenderFuncs is a unit test that verifies the output of each rendering function.
func TestRenderFuncs(t *testing.T) {
	// Define a test case struct that contains the necessary information to run the test.
	type testCase struct {
		test           string          // Name of the test case.
		kind           item.KindOfFile // Name of the rendering function to test.
		outputTestName string          // Path to the file that contains the expected output.
	}
	// Define the test cases to run.
	testCases := []testCase{
		{
			test:           "main render",
			kind:           item.Main,
			outputTestName: render.MainOutputPath,
		},
		{
			test:           "test render",
			kind:           item.Test,
			outputTestName: render.TestOutputPath,
		},
		{
			test:           "obj render",
			kind:           item.Object,
			outputTestName: render.ObjectOutputPath,
		},
		{
			test:           "tdd render",
			kind:           item.TableDriven,
			outputTestName: render.TddOutputPath,
		},
	}

	// exampleModel is an instance of the Model struct that can be used for testing.
	exampleModel := model{
		ObjectName: "Entry",
		Package:    "package entry",
	}

	require.NoError(t, os.Chdir("../../cmd"))

	// Iterate over each test case and run the test.
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			buffer := "buffer"
			_, errCreating := os.Create(buffer)
			require.NoError(t, errCreating)

			// Read the expected output from the file system.
			bytesContent, errRead := os.ReadFile(tc.outputTestName)
			require.NoError(t, errRead)

			require.NoError(t, render.Funcs[tc.kind](buffer, exampleModel))
			bytesExpected, errRead := os.ReadFile(buffer)
			require.NoError(t, errRead)

			// Verify that the output matches the expected output.
			require.Equal(t, bytesContent, bytesExpected)

			// Clean the buffer
			require.NoError(t, os.Remove(buffer))
		})
	}
}

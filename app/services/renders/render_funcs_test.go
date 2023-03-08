package renders

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Model represents the data that will be used to render the templates.
type Model struct {
	ObjectName string // Name of the main object in the Go file.
	Package    string // Name of the package the Go file belongs to.
}

// exampleModel is an instance of the Model struct that can be used for testing.
var exampleModel = Model{
	ObjectName: "Entry",
	Package:    "entry",
}

// TestRenderFuncs is a unit test that verifies the output of each rendering function.
func TestRenderFuncs(t *testing.T) {
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
			outputTestName: mainOutputPath,
		},
		{
			test:           "test render",
			kind:           "test",
			outputTestName: testOutputPath,
		},
		// {
		// 	test:           "obj render",
		// 	kind:           "object",
		// 	outputTestName: objectOutputPath,
		// },
		// {
		// 	test:           "tdd render",
		// 	kind:           "tableDriven",
		// 	outputTestName: tddOutputPath,
		// },
	}

	// Iterate over each test case and run the test.
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			// Render the template and capture the output.
			var buffer bytes.Buffer
			require.NoError(t, RenderFuncs[tc.kind](&buffer, exampleModel))

			// Read the expected output from the file system.
			bytesContent, errRead := os.ReadFile(tc.outputTestName)
			require.NoError(t, errRead)

			log.Print(string(bytesContent))
			log.Print(buffer.String())

			b1, _ := os.Create("exp.tmp")
			b2, _ := os.Create("act.tmp")

			b1.Write(bytesContent)
			b2.Write(buffer.Bytes())
			// Verify that the output matches the expected output.
			require.Equal(t, bytesContent, buffer.Bytes())

			//os.Remove("b1")
			//os.Remove("b2")
		})
	}
}

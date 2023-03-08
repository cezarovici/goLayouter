package renders

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type Model struct {
	FileName   string
	ObjectName string
	Package    string
}

var exampleModel = Model{
	FileName:   "entry.go",
	ObjectName: "Entry",
	Package:    "entry",
}

func TestRenderFuncs(t *testing.T) {
	type testCase struct {
		test           string
		kind           string
		outputTestName string
	}

	testCases := []testCase{
		{
			test:           "main render",
			kind:           "main",
			outputTestName: _mainOutput,
		},
		{
			test:           "test render",
			kind:           "test",
			outputTestName: _testOutput,
		},
		{
			test:           "obj render",
			kind:           "object",
			outputTestName: _objectOutput,
		},
		{
			test:           "tdd render",
			kind:           "tableDriven",
			outputTestName: _tddOutput,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			{
				var buffer bytes.Buffer
				require.NoError(t, RenderFuncs[tc.kind](&buffer, exampleModel))

				bytesContent, errRead := os.ReadFile(tc.outputTestName)
				require.NoError(t, errRead)

				require.Equal(t, bytesContent, buffer.Bytes())
			}
		})
	}
}

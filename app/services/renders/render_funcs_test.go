package renders

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/udhos/equalfile"
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
				buffer, errCreate := os.Create(_buffer)
				require.NoError(t, errCreate)

				require.NoError(t, RenderFuncs[tc.kind](buffer, exampleModel))

				cmp := equalfile.New(nil, equalfile.Options{}) // compare using single mode
				equal, errCompare := cmp.CompareFile(buffer.Name(), tc.outputTestName)

				require.NoError(t, buffer.Close())
				require.NoError(t, errCompare)
				require.NoError(t, os.Remove(buffer.Name()))

				require.Equal(t, true, equal)
			}
		})
	}
}

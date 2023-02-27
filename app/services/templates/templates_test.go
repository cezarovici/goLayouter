package templates

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/udhos/equalfile"
)

const (
	_testCases = "templateCases/"

	_mainOutput   = _testCases + "main/main_output.go"
	_objectOutput = _testCases + "object/object_output.go"
	_testOutput   = _testCases + "test/test_output.go"
	_tddOutput    = _testCases + "tdd/tdd_output.go"

	_mainInput   = _testCases + "main/main_input"
	_objectInput = _testCases + "object/object_input"
	_testInput   = _testCases + "test/test_input"
	_tddInput    = _testCases + "tdd/tableDriven_input"
)

type Model struct {
	FileName   string
	ObjectName string
	Package    string
}

var m = Model{
	FileName:   "entry.go",
	ObjectName: "Entry",
	Package:    "entry",
}

func TestRenderToPath(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		output      string
	}{
		{
			description: "main template",
			input:       _mainInput,
			output:      _mainOutput,
		},
		// {
		// 	description: "test template",
		// 	input:       _testInput,
		// 	output:      _testOutput,
		// },
		// {
		// 	description: "tdd template",
		// 	input:       _tddInput,
		// 	output:      _tddOutput,
		// },
		// {
		// 	description: "object template",
		// 	input:       _objectInput,
		// 	output:      _objectOutput,
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			{
				buffer, errCreate := os.Create("buffer.go")
				require.NoError(t, errCreate)
				require.NoError(t, RenderTo(buffer, tc.input, m))

				cmp := equalfile.New(nil, equalfile.Options{}) // compare using single mode
				equal, errCompare := cmp.CompareFile(buffer.Name(), tc.output)

				require.NoError(t, errCompare)
				require.Equal(t, true, equal)
				require.NoError(t, os.Remove(buffer.Name()))
			}
		})
	}
}

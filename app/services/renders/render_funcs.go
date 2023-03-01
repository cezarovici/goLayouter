package renders

import (
	"html/template"
	"io"
)

const (
	_testCases = "../templates/"

	_input  = "input/"
	_output = "output/"

	_mainOutput   = _testCases + _output + "main"
	_objectOutput = _testCases + _output + "object"
	_testOutput   = _testCases + _output + "test"
	_tddOutput    = _testCases + _output + "tdd"

	_mainIntput  = _testCases + _input + "main"
	_objectInput = _testCases + _input + "object"
	_testInput   = _testCases + _input + "test"
	_tddInput    = _testCases + _input + "tdd"

	_buffer = "buffer"
)

func renderTo(w io.Writer, templateFilePath string, model any) error {
	t, errParse := template.ParseFiles(templateFilePath)
	if errParse != nil {
		return errParse
	}

	return t.Execute(w, model)
}

var RenderFuncs = map[string]func(io.Writer, any) error{
	"main":        renderMain,
	"test":        renderTest,
	"object":      renderObject,
	"tableDriven": renderTableDriven,
}

func renderMain(w io.Writer, object any) error {
	return renderTo(w, _mainIntput, object)
}

func renderTest(w io.Writer, object any) error {
	return renderTo(w, _testInput, object)
}

func renderObject(w io.Writer, object any) error {
	return renderTo(w, _objectInput, object)
}

func renderTableDriven(w io.Writer, object any) error {
	return renderTo(w, _tddInput, object)
}

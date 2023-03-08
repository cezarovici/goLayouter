package renders

import (
	"io"
	"os"
	"text/template"
)

func renderTo(w io.Writer, templateFilePath string, model any) error {
	if _, err := os.Stat(templateFilePath); os.IsNotExist(err) {
		return err
	}

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

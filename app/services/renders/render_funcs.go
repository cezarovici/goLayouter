// Package renders provides rendering functionality for generating Go code files using templates.
package renders

import (
	"io"
	"log"
	"os"
	"text/template"
)

// renderTo renders the given model data to the specified template file and writes the output to the provided writer.
func renderTo(w io.Writer, templateFilePath string, model any) error {
	// Check if the specified template file exists.
	if _, err := os.Stat(templateFilePath); os.IsNotExist(err) {
		return err
	}

	// Parse the template file.
	t, errParse := template.ParseFiles(templateFilePath)
	if errParse != nil {
		return errParse
	}
	log.Print(model)
	// Execute the template with the model data and write the output to the writer.
	return t.Execute(w, model)
}

// RenderFuncs is a map of render function names to their corresponding functions.
// Each function takes a writer and a model as input, and generates Go code using the corresponding template.
var RenderFuncs = map[string]func(io.Writer, any) error{
	"main":        renderMain,
	"test":        renderTest,
	"object":      renderObject,
	"tableDriven": renderTableDriven,
}

// renderMain is a render function to render a main file from a template
func renderMain(w io.Writer, object any) error {
	return renderTo(w, mainInputPath, object)
}

// renderTest is a render function to render a test file from a template
func renderTest(w io.Writer, object any) error {
	return renderTo(w, testInputPath, object)
}

// renderObject is a render function to render an object file from a template
func renderObject(w io.Writer, object any) error {
	return renderTo(w, objectInputPath, object)
}

// renderTableDriven is a render function to render a table-driven test file from a template
func renderTableDriven(w io.Writer, object any) error {
	return renderTo(w, tddInputPath, object)
}

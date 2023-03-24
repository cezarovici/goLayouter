// Package renders provides rendering functionality for generating
// Go code files using templates.
package render

import (
	"io/fs"
	"os"
	"text/template"

	apperrors "github.com/cezarovici/goLayouter/app/errors"
)

const (
	o_rdwr fs.FileMode = 0o755
)

// renderTo renders the given model data to the specified template file and
// writes the output to the provided writer.
func renderTo(renderToPath string, templateFilePath string, model any) error {
	// Check if the specified template file exists.
	if _, err := os.Stat(templateFilePath); os.IsNotExist(err) {
		return &apperrors.RenderError{
			Caller:     "Renders",
			MethodName: "os.Stat",
			Issue:      err,
		}
	}

	// Parse the template file.
	templ, errParse := template.ParseFiles(templateFilePath)
	if errParse != nil {
		return &apperrors.RenderError{
			Caller:     "Renders",
			MethodName: "template parse files",
			Issue:      errParse,
		}
	}

	file, errCreate := os.OpenFile(renderToPath, os.O_RDWR, o_rdwr)
	if errCreate != nil {
		return &apperrors.RenderError{
			Caller:     "Renders",
			MethodName: "os open file",
			Issue:      errCreate,
		}
	}
	defer file.Close()

	// Execute the template with the model data and write the output to the writer.
	errTemplating := templ.Execute(file, model)
	if errTemplating != nil {
		return &apperrors.RenderError{
			Caller:     "Renders",
			MethodName: "template execute",
			Issue:      errTemplating,
		}
	}

	return nil
}

// renderMain is a render function to render a main file from a template.
func renderMain(renderToPath string, object any) error {
	return renderTo(renderToPath, MainInputPath, object)
}

// renderTest is a render function to render a test file from a template
func renderTest(renderToPath string, object any) error {
	return renderTo(renderToPath, TestInputPath, object)
}

// renderObject is a render function to render an object file from a template
func renderObject(renderToPath string, object any) error {
	return renderTo(renderToPath, ObjectInputPath, object)
}

// renderTableDriven is a render function
// to render a table-driven test file from a template
func renderTableDriven(renderToPath string, object any) error {
	return renderTo(renderToPath, TddInputPath, object)
}

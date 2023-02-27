package templates

import (
	"html/template"
	"io"
)

func RenderTo(w io.Writer, templateFilePath string, model any) error {
	t, errParse := template.ParseFiles(templateFilePath)
	if errParse != nil {
		return errParse
	}

	return t.Execute(w, model)
}

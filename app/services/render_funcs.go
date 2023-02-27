package services

import (
	"io"

	"github.com/cezarovici/goLayouter/app/services/templates"
)

const _templatesPath = "templates/input/"

const (
	_templateMain        = _templatesPath + "main"
	_templateTest        = _templatesPath + "test"
	_templateObject      = _templatesPath + "object"
	_templateTableDriven = _templatesPath + "tableDriven"
)

var _renderFuncs = map[string]func(io.Writer, any) error{
	"main":        renderMain,
	"test":        renderTest,
	"object":      renderObject,
	"tableDriven": renderTableDriven,
}

func renderMain(w io.Writer, object any) error {
	return templates.RenderTo(w, _templateMain, object)
}

func renderTest(w io.Writer, object any) error {
	return templates.RenderTo(w, _templateTest, object)
}

func renderObject(w io.Writer, object any) error {
	return templates.RenderTo(w, _templateObject, object)
}

func renderTableDriven(w io.Writer, object any) error {
	return templates.RenderTo(w, _templateTableDriven, object)
}

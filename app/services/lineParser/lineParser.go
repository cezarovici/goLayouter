package lineparser

import (
	"strings"

	"github.com/cezarovici/goLayouter/helpers"
)

// case 1
// file.go || folder -> ignoring
/// DONE

// case 2
// obj_myObject -> template object ( simple case )
// with file name without obj
// DONE

// case 3
// func_test.go -> template test //TODO -> check if there is any objects

// case 4
// obj_myObj_test.go ->

type lineParser struct {
	ObjectName string
	Package    string
}

func (l *lineParser) SetObjectName(path string) {
	(*l).ObjectName = strings.Title(helpers.GetLastPath(path))
}

func (l *lineParser) SetPackage(path string) {
	(*l).Package = helpers.GetLastPath(path)
}

package lineparser

import (
	"path"

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
//DONE

// case 4
// obj_myObj_test.go ->

type lineParser struct {
	ObjectName string
	Package    string
}

func (l *lineParser) SetObjectName(filePath string) {
	var (
		objName  string
		fileName string
	)

	_, fileName = path.Split(filePath)

	switch helpers.KindOfFile(fileName) {
	case "test", "object":
		objName = helpers.ExtractObjectFrom(fileName)

	default:
		objName = ""
	}

	(*l).ObjectName = objName
}

func (l *lineParser) SetPackage(path string) {
	(*l).Package = helpers.GetPackageFrom(path)
}

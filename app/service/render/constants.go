package render

import "github.com/cezarovici/goLayouter/domain/item"

// Funcs is a map of render function names to their corresponding functions.
// Each function takes a writer and a model as input,
// and generates Go code using the corresponding template.
var Funcs = map[item.KindOfFile]func(string, any) error{
	item.Main:        renderMain,
	item.Test:        renderTest,
	item.Object:      renderObject,
	item.TableDriven: renderTableDriven,
}

const (
	// Path from main to run templates:
	pathFromMain = "../service/templates/"
)

// Input and output directories:
const (
	inputDir  = "input/"
	outputDir = "output/"
)

// Output file paths:
const (
	MainOutputPath   = pathFromMain + outputDir + "main_result"
	ObjectOutputPath = pathFromMain + outputDir + "object_result"
	TestOutputPath   = pathFromMain + outputDir + "test_result"
	TddOutputPath    = pathFromMain + outputDir + "tableDriven_result"
)

// Input file paths
const (
	MainInputPath   = pathFromMain + inputDir + "main.gotmpl"
	ObjectInputPath = pathFromMain + inputDir + "object.gotmpl"
	TestInputPath   = pathFromMain + inputDir + "test.gotmpl"
	TddInputPath    = pathFromMain + inputDir + "tableDriven.gotmpl"
)

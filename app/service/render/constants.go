package render

const (
	// Path from main to run templates
	pathFromMain = "../service/templates/"
)

// Input and output directories
const (
	inputDir  = "input/"
	outputDir = "output/"
)

// Output file paths
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

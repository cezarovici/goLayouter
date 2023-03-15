package renders

const (
	// Path from main to run templates
	pathFromMain = "../services/templates/"
)

// Input and output directories
const (
	inputDir  = "input/"
	outputDir = "output/"
)

// Output file paths
const (
	_mainOutputPath   = pathFromMain + outputDir + "main_result"
	_objectOutputPath = pathFromMain + outputDir + "object_result"
	_testOutputPath   = pathFromMain + outputDir + "test_result"
	_tddOutputPath    = pathFromMain + outputDir + "tableDriven_result"
)

// Input file paths
const (
	_mainInputPath   = pathFromMain + inputDir + "main.gotmpl"
	_objectInputPath = pathFromMain + inputDir + "object.gotmpl"
	_testInputPath   = pathFromMain + inputDir + "test.gotmpl"
	_tddInputPath    = pathFromMain + inputDir + "tableDriven.gotmpl"
)

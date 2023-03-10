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
	mainOutputPath   = pathFromMain + outputDir + "main_result"
	objectOutputPath = pathFromMain + outputDir + "object_result"
	testOutputPath   = pathFromMain + outputDir + "test_result"
	tddOutputPath    = pathFromMain + outputDir + "tableDriven_result"
)

// Input file paths
const (
	mainInputPath   = pathFromMain + inputDir + "main.gotmpl"
	objectInputPath = pathFromMain + inputDir + "object.gotmpl"
	testInputPath   = pathFromMain + inputDir + "test.gotmpl"
	tddInputPath    = pathFromMain + inputDir + "tableDriven.gotmpl"
)

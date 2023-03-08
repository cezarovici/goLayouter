package renders

// Test cases directory
const (
	testCasesDir = "../templates/"
)

// Input and output directories
const (
	inputDir  = "input/"
	outputDir = "output/"
)

// Output file paths
const (
	mainOutputPath   = testCasesDir + outputDir + "main_result"
	objectOutputPath = testCasesDir + outputDir + "object_result"
	testOutputPath   = testCasesDir + outputDir + "test_result"
	tddOutputPath    = testCasesDir + outputDir + "tableDriven_result"
)

// Input file paths
const (
	mainInputPath   = testCasesDir + inputDir + "main.gotmpl"
	objectInputPath = testCasesDir + inputDir + "object.gotmpl"
	testInputPath   = testCasesDir + inputDir + "test.gotmpl"
	tddInputPath    = testCasesDir + inputDir + "tableDriven.gotmpl"
)

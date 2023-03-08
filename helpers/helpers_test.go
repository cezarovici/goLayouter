package helpers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

const testCasesPath = "../testCases/readTest/"

// ReadFile reads a file and returns its content as an array of strings
// It takes a filepath as input and returns a slice of strings with the file's content
// If there's any error, it returns an error
func TestReadFile(t *testing.T) {
	type testCase struct {
		test   string
		input  string
		output []string

		errorExpected error
	}

	testCases := []testCase{
		// Cases that generates erros
		{
			test:          "empty file",
			input:         testCasesPath + "emptyFile",
			output:        nil,
			errorExpected: errors.New("empty file passed"),
		},
		{
			test:          "no valid file",
			input:         testCasesPath + "invalid path",
			output:        nil,
			errorExpected: errors.New("not a valid file parsed"),
		},

		// Valid cases
		{
			test:          "1 line",
			input:         testCasesPath + "oneLine",
			output:        []string{"This is a line"},
			errorExpected: nil,
		},
		{
			test:          "Multiple lines",
			input:         testCasesPath + "multipleLines",
			output:        []string{"Line 1", "Line 2", "Line 3"},
			errorExpected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			output, errorReading := ReadFile(tc.input)

			require.Equal(t, tc.errorExpected, errorReading)
			require.Equal(t, output, tc.output)
		})
	}

}

// TypeOfFile takes a file path and returns a string indicating its type
// It takes a filepath as input and returns a string with the file's type
func TestTypeOfFile(t *testing.T) {
	type testCase struct {
		test   string
		input  string
		output string
	}

	testCases := []testCase{
		{
			test:   "file input",
			input:  "golangFile.go",
			output: "file",
		},
		{
			test:   "path input",
			input:  "! /home/user/ram",
			output: "path",
		},
		{
			test:   "package input",
			input:  "# package main",
			output: "package",
		},
		{
			test:   "folder input",
			input:  "folderName",
			output: "folder",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, TypeOfFile(tc.input))
		})
	}
}

// ToCurentDirectory checks if a file path is in the current directory or not
// It takes a filepath as input and returns a boolean indicating whether
// the path is in the current directory or not
func TestToCurentDirectory(t *testing.T) {
	type testCase struct {
		test   string
		input  string
		output bool
	}

	testCases := []testCase{
		{
			test:   "path to curent directory",
			input:  "! .",
			output: true,
		},
		{
			test:   "path to another directory",
			input:  "! /home/user/ram",
			output: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, ToCurentDirectory(tc.input))
		})
	}
}

// TestRemoveSelector is a unit test function for the RemoveSelector function.
// It tests the function's ability to remove a selector from a given string input.
func TestRemoveSelector(t *testing.T) {
	type testCase struct {
		test   string
		input  string
		output string
	}

	testCases := []testCase{
		{
			test:   "path splitted",
			input:  "! document",
			output: "document",
		},
		{
			test:   "package splitted",
			input:  "# package main",
			output: "package main",
		},
		{
			test:   "test package",
			input:  "# t",
			output: "t",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, RemoveSelector(tc.input))
		})
	}
}

// TestKindOfFile is a unit test function to test the KindOfFile function
// It tests the KindOfFile function with different inputs and expected
// outputs using a table-driven approach.
func TestKindOfFile(t *testing.T) {
	type testCase struct {
		test   string
		input  string
		output string
	}

	testCases := []testCase{
		{
			test:   "test main",
			input:  "main.go",
			output: "main",
		},
		{
			test:   "test file",
			input:  "func_test.go",
			output: "test",
		},
		{
			test:   "object file",
			input:  "obj_file.go",
			output: "object",
		},
		{
			test:   "object test",
			input:  "obj_file_test.go",
			output: "test",
		},
		{
			test:   "normal file",
			input:  "file.go",
			output: "normalFile",
		},
		{
			test:   "input folder",
			input:  "folder1",
			output: "folder",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, KindOfFile(tc.input))
		})
	}
}

// TestIsTestPackage is a unit test function to verify the behavior of the IsTestPackage function.
// It takes in a list of test cases, where each test case consists of an input string representing a package name and
// an expected output boolean indicating whether the package is a test package or not.
func TestIsTestPackage(t *testing.T) {
	type testCase struct {
		test   string
		input  string
		output bool
	}

	testCases := []testCase{
		{
			test:   "package test type 1",
			input:  "t",
			output: true,
		},
		{
			test:   "package test type 2",
			input:  "tt",
			output: true,
		},
		{
			test:   "not a test package",
			input:  "main",
			output: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, IsTestPackage(tc.input))
		})
	}
}

// TestCreatingGolangTestFile is a test function that tests the functionality of
// the CreateGolangTestFile function. It tests whether the function creates a new test file with the
// correct name based on the input file name.
func TestCreatingGolangTestFile(t *testing.T) {
	type testCase struct {
		test  string
		input string

		checkResult func(string, error)
	}

	testCases := []testCase{
		{
			test:  "just 1 file",
			input: "main.go",
			checkResult: func(s string, err error) {
				require.NoError(t, err)
				require.Equal(t, "main_test.go", s)
			},
		},
		{
			test:  "file from previous path",
			input: "../main.go",
			checkResult: func(s string, err error) {
				require.NoError(t, err)
				require.Equal(t, "../main_test.go", s)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			tc.checkResult(CreateGolangTestFile(tc.input))
		})
	}
}

// This function TestSplitline tests the SplitLine function.
// The SplitLine function takes a string representing a list of filenames
// separated by spaces and converts
// them to a list of filenames and test filenames, based on the package name provided.
func TestSplitline(t *testing.T) {
	type testCase struct {
		test        string
		input       string
		packageName string
		output      []string
	}

	testCases := []testCase{
		{
			test:        "converting test file with package type 1",
			input:       "file.go main.go head.go",
			packageName: "t",
			output:      []string{"file.go", "file_test.go", "main.go", "main_test.go", "head.go", "head_test.go"},
		},
		{
			test:        "converting test file with package type 2",
			input:       "file.go main.go head.go",
			packageName: "tt",
			output:      []string{"file.go", "file_test.go", "main.go", "main_test.go", "head.go", "head_test.go"},
		},
		{
			test:        "convert non test files",
			input:       "file.go main.go head.go",
			packageName: "main",
			output:      []string{"file.go", "main.go", "head.go"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, SplitLine(tc.input, tc.packageName))
		})
	}
}

// GetLastPath returns the last folder in a given path as a package declaration string.
// It takes a string as input representing the path to a folder.
// If the input is an empty string, it returns "package main".
// Otherwise, it splits the input path by the "/" separator and returns the package name declared in the last folder.
// For example, given the input "folder/subfolder1/subsubfolder1", the function returns "package subsubfolder1".
func TestGetLastPath(t *testing.T) {
	type testCase struct {
		test   string
		input  string
		output string
	}

	testCases := []testCase{
		{
			test:   "empty root",
			input:  "",
			output: "package main",
		},
		{
			test:   "just 1 file",
			input:  "folder",
			output: "package folder",
		},
		{
			test:   "2 folders",
			input:  "folder/folder1",
			output: "package folder1",
		},
		{
			test:   "mutitple folders",
			input:  "folder/subfolder1/subsubfolder1",
			output: "package subsubfolder1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, GetLastPath(tc.input))
		})
	}
}

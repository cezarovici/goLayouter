package helpers_test

import (
	"testing"

	"github.com/cezarovici/goLayouter/domain/item"
	"github.com/cezarovici/goLayouter/helpers"
	"github.com/stretchr/testify/require"
)

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
			test:   "path to current directory",
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
			require.Equal(t, tc.output, helpers.ToCurentDirectory(tc.input))
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
			require.Equal(t, tc.output, helpers.RemoveSelector(tc.input))
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
		output item.KindOfFile
	}

	testCases := []testCase{
		{
			test:   "test main",
			input:  "main.go",
			output: item.Main,
		},
		{
			test:   "test file",
			input:  "func_test.go",
			output: item.Test,
		},
		{
			test:   "object file",
			input:  "obj_file.go",
			output: item.Object,
		},
		{
			test:   "object test",
			input:  "obj_file_test.go",
			output: item.Test,
		},
		{
			test:   "normal file",
			input:  "file.go",
			output: item.NormalFile,
		},
		{
			test:   "input folder",
			input:  "folder1",
			output: item.Folder,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, helpers.KindOfFile(tc.input))
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
			require.Equal(t, tc.output, helpers.IsTestPackage(tc.input))
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
			tc.checkResult(helpers.CreateGolangTestFile(tc.input))
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
			require.Equal(t, tc.output, helpers.SplitLine(tc.input, tc.packageName))
		})
	}
}

// GetLastPath returns the last folder in a given path as a package declaration string.
// It takes a string as input representing the path to a folder.
// If the input is an empty string, it returns "package main".
// Otherwise, it splits the input path by the "/" separator and returns the package name declared in the last folder.
// For example, given the input "folder/subfolder1/subsubfolder1", the function returns "package subsubfolder1".
func TestGetPackageFromPath(t *testing.T) {
	t.Parallel()

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
			t.Parallel()
			require.Equal(t, tc.output, helpers.GetPackageFrom(tc.input))
		})
	}
}

func TestRemoveObjectKey(t *testing.T) {
	t.Parallel()
	type testCase struct {
		test   string
		input  string
		output string
	}

	testCases := []testCase{
		{
			test:   "exception 1",
			input:  "obj.go",
			output: "obj.go",
		},
		{
			test:   "exception 2",
			input:  "file.go",
			output: "file.go",
		},
		{
			test:   "case 1",
			input:  "obj_myObj.go",
			output: "myObj.go",
		},
		{
			test:   "case 2",
			input:  "object_myObj.go",
			output: "myObj.go",
		},
		{
			test:   "case 3",
			input:  "obj_test_myObj.go",
			output: "test_myObj.go",
		},
		{
			test:   "case 4",
			input:  "file_test.go",
			output: "file_test.go",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.output, helpers.RemoveObjectPrefix(tc.input))
		})
	}
}

func TestExtractObjectFrom(t *testing.T) {
	type testCase struct {
		test   string
		input  string
		output string
	}

	testCases := []testCase{
		{
			test:   "simple object file",
			input:  "obj_file.go",
			output: "File",
		},
		{
			test:   "test object file",
			input:  "obj_file_test.go",
			output: "File",
		},
		{
			test:   "simple file",
			input:  "file.go",
			output: "",
		},
		{
			test:   "without _",
			input:  "objectFile.go",
			output: "File",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.output, helpers.ExtractObjectFrom(tc.input))
		})
	}
}

func TestConvertToObjectName(t *testing.T) {
	t.Parallel()
	type testCase struct {
		test   string
		input  string
		output string
	}

	testCases := []testCase{
		{
			test:   "simple object file",
			input:  "app/test/obj_file.go",
			output: "File",
		},
		{
			test:   "simple test file",
			input:  "app/test/file_test.go",
			output: "",
		},
		{
			test:   "obj test file",
			input:  "app/test/obj_file_test.go",
			output: "File",
		},
		{
			test:   "unusual",
			input:  "obj.go",
			output: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.output, helpers.ConvertToObjectName(tc.input))
		})
	}
}

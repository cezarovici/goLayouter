package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const testCasesPath = "../testCases/readTest/"

// ReadFile reads a file and returns its content as an array of strings
// It takes a filepath as input and returns a slice of strings with the file's content
// If there's any error, it returns an error.
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
			errorExpected: fmt.Errorf("empty file passed"),
		},
		{
			test:          "no valid file",
			input:         testCasesPath + "invalid path",
			output:        nil,
			errorExpected: fmt.Errorf("not a valid file parsed"),
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

	for _, currentTestCase := range testCases {
		currentTestCase := currentTestCase

		t.Run(currentTestCase.test, func(t *testing.T) {
			t.Parallel()

			output, errorReading := ReadFile(currentTestCase.input)

			require.Equal(t, currentTestCase.errorExpected, errorReading)
			require.Equal(t, output, currentTestCase.output)
		})
	}
}

// TypeOfFile takes a file path and returns a string indicating its type
// It takes a filepath as input and returns a string with the file's type.
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

	for _, currentTestCase := range testCases {
		currentTestCase := currentTestCase

		t.Run(currentTestCase.test, func(t *testing.T) {
			require.Equal(t, currentTestCase.output, TypeOf(currentTestCase.input))
		})
	}
}

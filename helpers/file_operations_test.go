package helpers_test

import (
	"errors"
	"testing"

	"github.com/cezarovici/goLayouter/helpers"
	"github.com/stretchr/testify/require"
)

const testCasesPath = "../testCases/readTest/"

// ReadFile reads a file and returns its content as an array of strings
// It takes a filepath as input and returns a slice of strings with the file's content
// If there's any error, it returns an error.
func TestReadFile(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			output, errorReading := helpers.ReadFile(tc.input)

			require.Equal(t, tc.errorExpected, errorReading)
			require.Equal(t, output, tc.output)
		})
	}
}

// TypeOfFile takes a file path and returns a string indicating its type
// It takes a filepath as input and returns a string with the file's type.
func TestTypeOfFile(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			require.Equal(t, tc.output, helpers.TypeOf(tc.input))
		})
	}
}

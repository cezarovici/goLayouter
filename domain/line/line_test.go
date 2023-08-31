package line

import (
	"fmt"
	"testing"

	"github.com/cezarovici/goLayouter/domain/file"
	"github.com/cezarovici/goLayouter/domain/folder"
	"github.com/cezarovici/goLayouter/domain/item"
	"github.com/cezarovici/goLayouter/helpers"
	"github.com/stretchr/testify/require"
)

func TestConvertToLine(t *testing.T) {
	type tescurrentTestCasease struct {
		test   string
		input  string
		output Line
	}

	tescurrentTestCaseases := []tescurrentTestCasease{
		{
			test:  "first line.line",
			input: "folder1",
			output: Line{
				Info:  "folder1",
				Level: 0,
			},
		},
		{
			test:  "different Level",
			input: "  subfolder",
			output: Line{
				Info:  "subfolder",
				Level: 2,
			},
		},
		{
			test:  "package",
			input: " # package",
			output: Line{
				Info:  "# package",
				Level: 1,
			},
		},
	}

	for _, currentTestCase := range tescurrentTestCaseases {
		currentTestCase := currentTestCase

		t.Run(currentTestCase.test, func(t *testing.T) {
			require.Equal(t, currentTestCase.output, ConvertToLine(currentTestCase.input))
		})
	}
}

func TestRemoveObjectKey(t *testing.T) {
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

	for _, currenttestCase := range testCases {
		currenttestCase := currenttestCase

		t.Run(currenttestCase.test, func(t *testing.T) {
			require.Equal(t, currenttestCase.output,
				helpers.RemoveObjectPrefix(currenttestCase.input))
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

	for _, currenttestCase := range testCases {
		currenttestCase := currenttestCase

		t.Run(currenttestCase.test, func(t *testing.T) {
			require.Equal(t, currenttestCase.output,
				ExtractObjectFrom(currenttestCase.input))
		})
	}
}

func TestConvertToObjectName(t *testing.T) {
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

	for _, currenttestCase := range testCases {
		currenttestCase := currenttestCase

		t.Run(currenttestCase.test, func(t *testing.T) {
			require.Equal(t, currenttestCase.output,
				ConvertToObjectName(currenttestCase.input))
		})
	}
}

func TestNewLines(t *testing.T) {
	type testCase struct {
		test   string
		input  []string
		output Lines

		errorExpected error
	}

	testCases := []testCase{
		// Cases with errors
		{
			test:          "no Package",
			input:         nil,
			output:        nil,
			errorExpected: fmt.Errorf("no Package parsed"),
		},

		// Happy cases
		{
			test:  "2 line.lines",
			input: []string{"folder1", " subfolder1"},
			output: Lines{
				Line{
					Info:  "folder1",
					Level: 0,
				},
				Line{
					Info:  "subfolder1",
					Level: 1,
				},
			},
			errorExpected: nil,
		},
	}

	for _, currentTestCase := range testCases {
		currentTestCase := currentTestCase

		t.Run(currentTestCase.test, func(t *testing.T) {
			lines, errCreatinLines := NewLines(currentTestCase.input)

			require.Equal(t, currentTestCase.errorExpected, errCreatinLines)
			require.Equal(t, currentTestCase.output, lines)
		})
	}
}

const _parseTescurrentTestCaseases = "../../testCases/parseTest/"

func TestToItems(t *testing.T) {
	type testCase struct {
		test          string
		input         string
		expectedItems *item.Items
	}

	testCases := []testCase{
		{
			test:  "folders with files",
			input: _parseTescurrentTestCaseases + "foldersWithFiles/input",
			expectedItems: &item.Items{
				item.Item{
					ObjectPath: folder.Folder{
						Path: "folder1",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: folder.Folder{
						Path: "folder1/subfolder1",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: file.File{
						Path:    "folder1/subfolder1/file.go",
						Package: "package subfolder1",
					},
					Kind: item.NormalFile,
				},
				item.Item{
					ObjectPath: file.File{
						Path:    "folder1/subfolder1/obj.go",
						Package: "package subfolder1",
					},
					Kind: item.Object,
				},
				item.Item{
					ObjectPath: file.File{
						Path:    "folder1/subfolder1/main.go",
						Package: "package main",
					},
					Kind: item.Main,
				},
				item.Item{
					ObjectPath: folder.Folder{
						Path: "folder2",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: file.File{
						Path:    "folder2/test1.go",
						Package: "package folder2",
					},
					Kind: item.Test,
				},
			},
		},
		{
			test:  "folder with indents",
			input: _parseTescurrentTestCaseases + "foldersWithIndents/input",
			expectedItems: &item.Items{
				item.Item{
					ObjectPath: folder.Folder{
						Path: "folder",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: folder.Folder{
						Path: "folder1",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: folder.Folder{
						Path: "folder1/subfolder1",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: folder.Folder{
						Path: "folder1/subfolder1/subsubfolder1",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: folder.Folder{
						Path: "folder1/subfolder2",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: folder.Folder{
						Path: "folder2",
					},
					Kind: item.Folder,
				},
			},
		},
		{
			test:  "folders with test packages",
			input: _parseTescurrentTestCaseases + "foldersWithTestPackage/input",
			expectedItems: &item.Items{
				item.Item{
					ObjectPath: folder.Folder{
						Path: "app",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: file.File{
						Path:    "app/main.go",
						Package: "package main",
					},
					Kind: item.Main,
				},
				item.Item{
					ObjectPath: folder.Folder{
						Path: "app/domain",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: file.File{
						Path:    "app/domain/interfaces.go",
						Package: "package domain",
					},
					Kind: item.NormalFile,
				},
				item.Item{
					ObjectPath: file.File{
						Path:    "app/domain/file.go",
						Package: "package file",
					},
					Kind: item.NormalFile,
				},
				item.Item{
					ObjectPath: file.File{
						Path:       "app/domain/file.go",
						Package:    "package file",
						ObjectName: "File",
					},
					Kind: item.Object,
				},
				item.Item{
					ObjectPath: file.File{
						Path:       "app/domain/file_test.go",
						Package:    "package file",
						ObjectName: "File",
					},
					Kind: item.Test,
				},
			},
		},
		{
			test:  "folders with change directory",
			input: _parseTescurrentTestCaseases + "foldersWithChangeDirectory/input",
			expectedItems: &item.Items{
				item.Item{
					ObjectPath: file.File{
						Path: "README.md",
					},
				},
				item.Item{
					ObjectPath: folder.Folder{
						Path: "temporary_folder",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: folder.Folder{
						Path: "temporary_folder/app",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: file.File{
						Path:    "temporary_folder/app/main.go",
						Package: "package main",
					},
					Kind: item.Main,
				},
				item.Item{
					ObjectPath: folder.Folder{
						Path: "temporary_folder/app/domain",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: file.File{
						Path:    "temporary_folder/app/domain/interfaces.go",
						Package: "package domain",
					},
					Kind: item.NormalFile,
				},
				item.Item{
					ObjectPath: folder.Folder{
						Path: "temporary_folder/app/domain/file",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: file.File{
						Path:    "temporary_folder/app/domain/file/file.go",
						Package: "package file",
					},
					Kind: item.NormalFile,
				},

				item.Item{
					ObjectPath: folder.Folder{
						Path: "temporary_folder/app/domain/obj",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: file.File{
						Path:       "temporary_folder/app/domain/obj/file.go",
						Package:    "package obj",
						ObjectName: "File",
					},
					Kind: item.Object,
				},
				item.Item{
					ObjectPath: folder.Folder{
						Path: "temporary_folder/student",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: file.File{
						Path:    "temporary_folder/student/student.go",
						Package: "package student",
					},
					Kind: item.NormalFile,
				},
				item.Item{
					ObjectPath: folder.Folder{
						Path: "temporary_folder/student/study",
					},
					Kind: item.Folder,
				},
				item.Item{
					ObjectPath: file.File{
						Path:    "temporary_folder/student/study/study.go",
						Package: "package study",
					},
					Kind: item.NormalFile,
				},
			},
		},
	}

	for _, currentTestCase := range testCases {
		currentTestCase := currentTestCase

		t.Run(currentTestCase.test, func(t *testing.T) {
			inputPackage, errorReading := helpers.ReadFile(currentTestCase.input)
			require.NoError(t, errorReading)

			lines, errNewLines := NewLines(inputPackage)
			require.NoError(t, errNewLines)

			require.Equal(t, currentTestCase.expectedItems, lines.ToItems())
		})
	}
}

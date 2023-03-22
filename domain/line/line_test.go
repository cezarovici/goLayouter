package line

import (
	"errors"
	"testing"

	"github.com/cezarovici/goLayouter/domain/file"
	"github.com/cezarovici/goLayouter/domain/folder"
	"github.com/cezarovici/goLayouter/domain/item"
	"github.com/cezarovici/goLayouter/helpers"
	"github.com/stretchr/testify/require"
)

func TestConvertToLine(t *testing.T) {
	type testCase struct {
		test   string
		input  string
		output Line
	}

	testCases := []testCase{
		{
			test:  "first line",
			input: "folder1",
			output: Line{
				info:  "folder1",
				level: 0,
			},
		},
		{
			test:  "different level",
			input: "  subfolder",
			output: Line{
				info:  "subfolder",
				level: 2,
			},
		},
		{
			test:  "package",
			input: " # package",
			output: Line{
				info:  "# package",
				level: 1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, ConvertToLine(tc.input))
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
			errorExpected: errors.New("no Package parsed"),
		},

		// Happy cases
		{
			test:  "2 lines",
			input: []string{"folder1", " subfolder1"},
			output: Lines{
				Line{
					info:  "folder1",
					level: 0,
				},
				Line{
					info:  "subfolder1",
					level: 1,
				},
			},
			errorExpected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			lines, errCreatingLines := NewLines(tc.input)

			require.Equal(t, tc.errorExpected, errCreatingLines)
			require.Equal(t, tc.output, lines)
		})
	}
}

const _parseTestCases = "../../testCases/parseTest/"

func TestToItems(t *testing.T) {
	type testCase struct {
		test          string
		input         string
		expectedItems *item.Items
	}

	testCases := []testCase{
		{
			test:  "folders with files",
			input: _parseTestCases + "foldersWithFiles/input",
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
			input: _parseTestCases + "foldersWithIndents/input",
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
			input: _parseTestCases + "foldersWithTestPackage/input",
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
			input: _parseTestCases + "foldersWithChangeDirectory/input",
			expectedItems: &item.Items{
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

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			inputPackage, errorReading := helpers.ReadFile(tc.input)
			require.NoError(t, errorReading)

			lines, errNewLines := NewLines(inputPackage)
			require.NoError(t, errNewLines)

			require.Equal(t, tc.expectedItems, lines.ToItems())
		})
	}

}

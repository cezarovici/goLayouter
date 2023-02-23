package line

import (
	"errors"
	"testing"

	"github.com/cezarovici/goLayouter/domain"
	"github.com/cezarovici/goLayouter/domain/file"
	"github.com/cezarovici/goLayouter/domain/folder"
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
			test:          "no content",
			input:         nil,
			output:        nil,
			errorExpected: errors.New("no content parsed"),
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

const _parseTestCases = "../testCases/parseTest/"

func TestToItems(t *testing.T) {
	type testCase struct {
		test          string
		input         string
		expectedItems *domain.Items
	}

	testCases := []testCase{
		// {
		// 	test:  "folders with files",
		// 	input: _parseTestCases + "foldersWithFiles/input",
		// 	expectedItems: &domain.Items{
		// 		domain.Item{
		// 			ObjectPath: folder.Folder{
		// 				Path: "folder1",
		// 			},
		// 			Kind: "folder",
		// 		},
		// 		domain.Item{
		// 			ObjectPath: folder.Folder{
		// 				Path: "folder1/subfolder1",
		// 			},
		// 			Kind: "folder",
		// 		},
		// 		domain.Item{
		// 			ObjectPath: file.File{
		// 				Path:    "folder1/subfolder1/file.go",
		// 				Content: "package subfolder1",
		// 			},
		// 			Kind: "normalFile",
		// 		},
		// 		domain.Item{
		// 			ObjectPath: file.File{
		// 				Path:    "folder1/subfolder1/obj.go",
		// 				Content: "package subfolder1",
		// 			},
		// 			Kind: "object",
		// 		},
		// 		domain.Item{
		// 			ObjectPath: file.File{
		// 				Path:    "folder1/subfolder1/main.go",
		// 				Content: "package main",
		// 			},
		// 			Kind: "main",
		// 		},
		// 		domain.Item{
		// 			ObjectPath: folder.Folder{
		// 				Path: "subfolder2",
		// 			},
		// 			Kind: "folder",
		// 		},
		// 		domain.Item{
		// 			ObjectPath: file.File{
		// 				Path:    "subfolder2/test1.go",
		// 				Content: "package subfolder2",
		// 			},
		// 			Kind: "test",
		// 		},
		// 	},
		// },
		// {
		// 	test:  "folder with indents",
		// 	input: _parseTestCases + "foldersWithIndents/input",
		// 	expectedItems: &domain.Items{
		// 		domain.Item{
		// 			ObjectPath: folder.Folder{
		// 				Path: "folder",
		// 			},
		// 			Kind: "folder",
		// 		},
		// 		domain.Item{
		// 			ObjectPath: folder.Folder{
		// 				Path: "folder1",
		// 			},
		// 			Kind: "folder",
		// 		},
		// 		domain.Item{
		// 			ObjectPath: folder.Folder{
		// 				Path: "folder1/subfolder1",
		// 			},
		// 			Kind: "folder",
		// 		},
		// 		domain.Item{
		// 			ObjectPath: folder.Folder{
		// 				Path: "folder1/subfolder1/subsubfolder1",
		// 			},
		// 			Kind: "folder",
		// 		},
		// 		domain.Item{
		// 			ObjectPath: folder.Folder{
		// 				Path: "folder1/subfolder2",
		// 			},
		// 			Kind: "folder",
		// 		},
		// 		domain.Item{
		// 			ObjectPath: folder.Folder{
		// 				Path: "folder2",
		// 			},
		// 			Kind: "folder",
		// 		},
		// 	},
		// },
		//{
		// 	test:  "folder with packages",
		// 	input: _parseTestCases + "foldersWithPackages/input",
		// 	expectedItems: &domain.Items{
		// 		domain.Item{
		// 			ObjectPath: folder.Folder{
		// 				Path: "person",
		// 			},
		// 			Kind: "folder",
		// 		},
		// 		domain.Item{
		// 			ObjectPath: file.File{
		// 				Path:    "person/person.go",
		// 				Content: "package person",
		// 			},
		// 			Kind: "normalFile",
		// 		},
		// 		domain.Item{
		// 			ObjectPath: folder.Folder{
		// 				Path: "person/student",
		// 			},
		// 			Kind: "folder",
		// 		},
		// 		domain.Item{
		// 			ObjectPath: file.File{
		// 				Path:    "person/student/student.go",
		// 				Content: "package student",
		// 			},
		// 			Kind: "normalFile",
		// 		},
		// 		domain.Item{
		// 			ObjectPath: file.File{
		// 				Path:    "person/student/study.go",
		// 				Content: "package studyInterests",
		// 			},
		// 			Kind: "normalFile",
		// 		},
		// 	},
		// },
		{
			test:  "folders with test packages",
			input: _parseTestCases + "foldersWithTestPackage/input",
			expectedItems: &domain.Items{
				domain.Item{
					ObjectPath: folder.Folder{
						Path: "app",
					},
					Kind: "folder",
				},
				domain.Item{
					ObjectPath: file.File{
						Path:    "app/main.go",
						Content: "package main",
					},
					Kind: "main",
				},
				domain.Item{
					ObjectPath: folder.Folder{
						Path: "app/domain",
					},
					Kind: "folder",
				},
				domain.Item{
					ObjectPath: file.File{
						Path:    "app/domain/interfaces.go",
						Content: "package domain",
					},
					Kind: "normalFile",
				},
				domain.Item{
					ObjectPath: file.File{
						Path:    "app/domain/file.go",
						Content: "package file",
					},
					Kind: "normalFile",
				},
				domain.Item{
					ObjectPath: file.File{
						Path:    "app/domain/obj_file.go",
						Content: "package file",
					},
					Kind: "object",
				},
				domain.Item{
					ObjectPath: file.File{
						Path:    "app/domain/obj_file_test.go",
						Content: "package file",
					},
					Kind: "test",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			inputContent, errorReading := helpers.ReadFile(tc.input)
			require.NoError(t, errorReading)

			lines, errNewLines := NewLines(inputContent)
			require.NoError(t, errNewLines)

			require.Equal(t, tc.expectedItems, lines.ToItems())
		})
	}

}

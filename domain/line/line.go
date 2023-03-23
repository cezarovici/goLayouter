package line

import (
	"fmt"
	"strings"

	"github.com/cezarovici/goLayouter/domain/file"
	"github.com/cezarovici/goLayouter/domain/folder"
	"github.com/cezarovici/goLayouter/domain/item"
	"github.com/cezarovici/goLayouter/helpers"
	"github.com/cezarovici/goLayouter/helpers/stack"
)

// Line represents a line of text and its indentation Level
type Line struct {
	Info  string // the text Package of the line
	Level int    // the indentation Level of the line
}

// Lines is a slice of Line
type Lines []Line

// Constants for default and test package types
const (
	_defaultPackage   = "package main" // default package
	_testPackageType1 = "# t"          // test package type 1
	_testPackageType2 = "# tt"         // test package type 2
)

// ConvertToLine converts a string to a Line struct, parsing the indentation Level
func ConvertToLine(line string) Line {
	// Trim any whitespace from the beginning of the line
	lineInfo := strings.TrimLeft(line, " ")

	// Calculate the indentation Level by subtracting the length of the trimmed text
	// from the length of the original line
	lineLevel := len(line) - len(lineInfo)

	return Line{
		Info:  lineInfo,
		Level: lineLevel,
	}
}

// NewLines converts a slice of strings to a slice of Lines,
// using ConvertToLine to parse each string
func NewLines(content []string) (Lines, error) {
	// Return an error if there is no Package to parse
	if len(content) == 0 {
		return nil, fmt.Errorf("no Package parsed")
	}

	// Initialize an empty slice of Lines
	var res Lines

	// Parse each string in the Package slice and append the resulting Line to the res slice
	for _, line := range content {
		res = append(res, ConvertToLine(line))
	}

	return res, nil
}

func (lines Lines) ToItems() *item.Items {
	// Create an empty slice of res
	var res item.Items

	// Set a flag to indicate if this is the first line being processed
	firstLine := true

	// Create stacks to keep track of the paths and indentation Levels
	var pathStack stack.Stack
	var indentStack stack.Stack

	// Initialize the package stack with the default package name
	packageStack := stack.Stack{_defaultPackage}

	// Iterate over each line in the input file
	for _, line := range lines {
		// Check the type of the line (empty, path, or file)
		switch helpers.TypeOf(line.Info) {
		case "empty":
			// If the line is empty, skip it
			continue

		case "path":
			// If the line is a path, reset the package stack, indent stack, and path stack
			packageStack = stack.Stack{_defaultPackage}
			indentStack = nil
			pathStack = nil

			// If the path is not the current directory,
			// add it to the path stack and create a new item
			if !helpers.ToCurentDirectory(line.Info) {
				pathStack.Push(helpers.RemoveSelector(line.Info))
				indentStack.Push(-1)

				// Create a new item with the path and kind of the current line
				res = append(res, item.Item{
					ObjectPath: folder.Folder{
						Path: pathStack.String(),
					},
					Kind: helpers.KindOfFile(line.Info),
				})

				// Set firstLine to false since we have already processed the first line
				firstLine = false

				// Continue to the next line
				continue
			}

			// If the path is the current directory,
			// push the current Level to the indent stack and continue
			indentStack.Push(line.Level)

			continue

		// Check if the type of the line is "package"
		case "package":
			// Push the package name onto the package stack
			packageStack.Push(helpers.RemoveSelector(line.Info))
			// Continue to the next line
			continue

		// If the current entry in the path stack is a file,
		// we need to determine the package name
		// based on the directory structure of the file path
		case "file":
			// Get the last path based on the directory structure of the file path
			packageName := helpers.GetPackageFrom(pathStack.String())
			// Check if the current package name is not the default package ("package main")
			if packageStack.Peek() != _defaultPackage {
				// If it's not the default package, use the current package name
				packageName = packageStack.Peek().(string)
			}

			// Check if the current package is a test package
			if helpers.IsTestPackage(packageStack.Peek().(string)) {
				// If it's a test package, peek the test package type (either "t" or "tt")
				testPackageType := packageStack.Peek()

				// Pop the test package type from the package stack to get the previous package
				packageStack.Pop()

				// Set the previous package name as the main package name
				packageName = packageStack.Peek().(string)

				// Push the test package type back onto the package stack
				packageStack.Push(testPackageType)
			}

			// Split the line into files based on the current package name
			files := helpers.SplitLine(line.Info, packageStack.Peek().(string))

			// Iterate over the files and determine the package name and kind for each file
			for _, fileName := range files {
				var objectName string
				isObject := false

				// If the file is a main package, use the default package name
				if helpers.KindOfFile(fileName) == item.Main {
					packageName = _defaultPackage
				}

				objectName = helpers.ConvertToObjectName(fileName)
				if helpers.KindOfFile(fileName) == item.Object || helpers.KindOfFile(fileName) == item.Test {
					fileName = helpers.RemoveObjectPrefix(fileName)

					if helpers.KindOfFile(fileName) != item.Test {
						isObject = true
					}
				}

				kind := helpers.KindOfFile(fileName)
				if isObject {
					kind = item.Object
				}

				// Create a new item with the file path and package name
				newFile := file.File{
					Path:       pathStack.String() + "/" + fileName,
					Package:    packageName,
					ObjectName: objectName,
				}

				newItem := item.Item{
					ObjectPath: newFile,
					Kind:       kind,
				}

				// Add the new item to the res list
				res = append(res, newItem)
			}

			continue
		}

		// If this is the first line of the file,
		// we need to push the file path onto the path stack
		// and push an initial indentation Level (0) onto the indent stack
		if firstLine {
			pathStack.Push(line.Info)
			indentStack.Push(0)

			// Create a new Folder object representing the directory containing the file
			folder := folder.Folder{
				Path: pathStack.String(),
			}

			// Add the Folder object to the list of res
			res = append(res, item.Item{
				ObjectPath: folder,
				Kind:       helpers.KindOfFile(line.Info),
			})

			// Set the firstLine flag to false so that this block of code is not executed again
			firstLine = false

			continue
		}

		// If the line Level is greater than the top Level on the stack, the line is a
		// subdirectory and should be pushed onto the stack. A new folder object should
		// be created and added to the res list.
		if line.Level > indentStack.Peek().(int) {
			pathStack.Push(line.Info)
			indentStack.Push(line.Level)

			folder := folder.Folder{
				Path: pathStack.String(),
			}
			res = append(res, item.Item{
				ObjectPath: folder,
				Kind:       helpers.KindOfFile(line.Info),
			})

			continue
		}

		// If the line Level is equal to the top Level on the stack, the line is in the
		// same directory as the previous line and should replace the previous line on
		// the stack. A new folder object should be created and added to the res list.
		if line.Level == indentStack.Peek().(int) {
			pathStack.Pop()

			pathStack.Push(line.Info)
			indentStack.Push(line.Level)

			res = append(res, item.Item{
				ObjectPath: folder.Folder{
					Path: pathStack.String(),
				},
				Kind: helpers.KindOfFile(line.Info),
			})

			continue
		}

		// If the line Level is less than the top Level on the stack, we need to
		// pop Items off the stack until we reach the parent directory of the current line.
		// Then, we can add the current line to the stack and create a new folder object
		// to be added to the Items list.
		for line.Level < indentStack.Peek().(int) && len(indentStack) > 1 {
			pathStack.Pop()
			indentStack.Pop()

			if line.Level == indentStack.Peek().(int) {
				pathStack.Pop()
				indentStack.Pop()

				break
			}
		}

		pathStack.Push(line.Info)
		indentStack.Push(line.Level)

		res = append(res, item.Item{
			ObjectPath: folder.Folder{
				Path: pathStack.String(),
			},
			Kind: helpers.KindOfFile(line.Info),
		})
	}

	return &res
}

// TODO add go.mod yaml.config github actions

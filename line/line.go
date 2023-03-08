package line

import (
	"fmt"
	"log"
	"strings"

	"github.com/cezarovici/goLayouter/domain/file"
	"github.com/cezarovici/goLayouter/domain/folder"
	"github.com/cezarovici/goLayouter/domain/item"
	"github.com/cezarovici/goLayouter/helpers"
	"github.com/cezarovici/goLayouter/helpers/stack"
)

// Line represents a line of text and its indentation level
type Line struct {
	info  string // the text content of the line
	level int    // the indentation level of the line
}

// Lines is a slice of Line
type Lines []Line

// Constants for default and test package types
const (
	_defaultPackage   = "package main" // default package
	_testPackageType1 = "# t"          // test package type 1
	_testPackageType2 = "# tt"         // test package type 2
)

// ConvertToLine converts a string to a Line struct, parsing the indentation level
func ConvertToLine(line string) Line {
	// Trim any whitespace from the beginning of the line
	lineInfo := strings.TrimLeft(line, " ")

	// Calculate the indentation level by subtracting the length of the trimmed text
	// from the length of the original line
	lineLevel := len(line) - len(lineInfo)

	return Line{
		info:  lineInfo,
		level: lineLevel,
	}
}

// NewLines converts a slice of strings to a slice of Lines, using ConvertToLine to parse each string
func NewLines(content []string) (Lines, error) {
	// Return an error if there is no content to parse
	if len(content) == 0 {
		return nil, fmt.Errorf("no content parsed")
	}

	// Initialize an empty slice of Lines
	var items Lines

	// Parse each string in the content slice and append the resulting Line to the items slice
	for _, line := range content {
		items = append(items, ConvertToLine(line))
	}

	return items, nil
}

func (lines Lines) ToItems() *item.Items {
	// Create an empty slice of items
	var items item.Items

	// Set a flag to indicate if this is the first line being processed
	firstLine := true

	// Create stacks to keep track of the paths and indentation levels
	var pathStack stack.Stack
	var indentStack stack.Stack

	// Initialize the package stack with the default package name
	packageStack := stack.Stack{_defaultPackage}

	// Iterate over each line in the input file
	for _, line := range lines {
		// Check the type of the line (empty, path, or file)
		switch helpers.TypeOfFile(line.info) {
		case "empty":
			// If the line is empty, skip it
			continue

		case "path":
			// If the line is a path, reset the package stack, indent stack, and path stack
			packageStack = stack.Stack{_defaultPackage}
			indentStack = nil
			pathStack = nil

			// If the path is not the current directory, add it to the path stack and create a new item
			if !helpers.ToCurentDirectory(line.info) {
				pathStack.Push(helpers.RemoveSelector(line.info))
				indentStack.Push(-1)

				// Create a new item with the path and kind of the current line
				items = append(items, item.Item{
					ObjectPath: folder.Folder{
						Path: pathStack.String(),
					},
					Kind: helpers.KindOfFile(line.info),
				})

				// Set firstLine to false since we have already processed the first line
				firstLine = false

				// Continue to the next line
				continue
			}

			// If the path is the current directory, push the current level to the indent stack and continue
			indentStack.Push(line.level)

			continue

		// Check if the type of the line is "package"
		case "package":
			// Push the package name onto the package stack
			packageStack.Push(helpers.RemoveSelector(line.info))
			// Continue to the next line
			continue

		// If the current entry in the path stack is a file, we need to determine the package name
		// based on the directory structure of the file path
		case "file":
			// Get the last path based on the directory structure of the file path
			packageName := helpers.GetLastPath(pathStack.String())
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
			files := helpers.SplitLine(line.info, packageStack.Peek().(string))

			// Iterate over the files and determine the package name and kind for each file
			for _, fileName := range files {
				isObject := false

				// If the file is a main package, use the default package name
				if helpers.KindOfFile(fileName) == "main" {
					packageName = _defaultPackage
				}

				if helpers.KindOfFile(fileName) == "object" || helpers.KindOfFile(fileName) == "test" {
					fileName = helpers.RemoveObjectKey(fileName)

					if helpers.KindOfFile(fileName) != "test" {
						isObject = true
					}
				}

				kind := helpers.KindOfFile(fileName)
				if isObject {
					kind = "object"
				}
				log.Print(fileName, " ", kind)

				// Create a new item with the file path and package name
				newFile := file.File{
					Path:    pathStack.String() + "/" + fileName,
					Content: packageName,
				}

				newItem := item.Item{
					ObjectPath: newFile,
					Kind:       kind,
				}

				// Add the new item to the items list
				items = append(items, newItem)
			}

			continue
		}

		// If this is the first line of the file, we need to push the file path onto the path stack
		// and push an initial indentation level (0) onto the indent stack
		if firstLine {
			pathStack.Push(line.info)
			indentStack.Push(0)

			// Create a new Folder object representing the directory containing the file
			folder := folder.Folder{
				Path: pathStack.String(),
			}

			// Add the Folder object to the list of items
			items = append(items, item.Item{
				ObjectPath: folder,
				Kind:       helpers.KindOfFile(line.info),
			})

			// Set the firstLine flag to false so that this block of code is not executed again
			firstLine = false

			continue
		}

		// If the line level is greater than the top level on the stack, the line is a
		// subdirectory and should be pushed onto the stack. A new folder object should
		// be created and added to the items list.
		if line.level > indentStack.Peek().(int) {
			pathStack.Push(line.info)
			indentStack.Push(line.level)

			folder := folder.Folder{
				Path: pathStack.String(),
			}
			items = append(items, item.Item{
				ObjectPath: folder,
				Kind:       helpers.KindOfFile(line.info),
			})

			continue
		}

		// If the line level is equal to the top level on the stack, the line is in the
		// same directory as the previous line and should replace the previous line on
		// the stack. A new folder object should be created and added to the items list.
		if line.level == indentStack.Peek().(int) {
			pathStack.Pop()

			pathStack.Push(line.info)
			indentStack.Push(line.level)

			items = append(items, item.Item{
				ObjectPath: folder.Folder{
					Path: pathStack.String(),
				},
				Kind: helpers.KindOfFile(line.info),
			})

			continue
		}

		// If the line level is less than the top level on the stack, we need to
		// pop items off the stack until we reach the parent directory of the current line.
		// Then, we can add the current line to the stack and create a new folder object
		// to be added to the items list.
		for line.level < indentStack.Peek().(int) && len(indentStack) > 1 {
			pathStack.Pop()
			indentStack.Pop()

			if line.level == indentStack.Peek().(int) {
				pathStack.Pop()
				indentStack.Pop()

				break
			}
		}

		pathStack.Push(line.info)
		indentStack.Push(line.level)

		items = append(items, item.Item{
			ObjectPath: folder.Folder{
				Path: pathStack.String(),
			},
			Kind: helpers.KindOfFile(line.info),
		})
	}

	return &items
}

//TODO add go.mod yaml.config github actions

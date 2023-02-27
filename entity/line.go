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

type Line struct {
	info  string
	level int
}

type Lines []Line

const (
	_defaultPackage   = "package main"
	_testPackageType1 = "# t"
	_testPackageType2 = "# tt"
)

func ConvertToLine(line string) Line {
	lineInfo := strings.TrimLeft(line, " ")

	lineLevel := len(line) - len(lineInfo)

	return Line{
		info:  lineInfo,
		level: lineLevel,
	}
}

func NewLines(content []string) (Lines, error) {
	if len(content) == 0 {
		return nil, fmt.Errorf("no content parsed")
	}

	var items Lines

	for _, line := range content {
		items = append(items, ConvertToLine(line))
	}

	return items, nil
}

func (lines Lines) ToItems() *item.Items {
	var items item.Items

	first := true

	var stackPaths stack.Stack
	var stackIndents stack.Stack
	stackPackages := stack.Stack{_defaultPackage}

	for _, line := range lines {
		switch helpers.TypeOfFile(line.info) {
		case "empty":
			// jump if is parsed and empty line
			continue

		case "path":
			stackPackages = stack.Stack{_defaultPackage}
			stackIndents = nil
			stackPaths = nil

			if !helpers.ToCurentDirectory(line.info) {
				stackPaths.Push(helpers.RemoveSelector(line.info))
				stackIndents.Push(-1)

				items = append(items, item.Item{
					ObjectPath: folder.Folder{
						Path: stackPaths.String(),
					},
					Kind: helpers.KindOfFile(line.info),
				})

				first = false

				continue
			}

			stackIndents.Push(line.level)

			continue

		case "package":
			stackPackages.Push(helpers.RemoveSelector(line.info))

			continue

		case "file":
			packageName := helpers.GetRootPackage(stackPaths.String())

			if stackPackages.Peek() != "package main" {
				packageName = stackPackages.Peek().(string)
			}

			if helpers.IsTestPackage(stackPackages.Peek().(string)) {
				testPackage := stackPackages.Peek()         // peek the test package ( t or tt )
				stackPackages.Pop()                         // poping to get the previous package
				packageName = stackPackages.Peek().(string) // setting the previous package

				stackPackages.Push(testPackage)
			}

			files := helpers.SplitLine(line.info, stackPackages.Peek().(string))
			for _, fileName := range files {
				if helpers.KindOfFile(fileName) == "main" {
					packageName = _defaultPackage
				}

				items = append(items, item.Item{
					ObjectPath: file.File{
						Path:    stackPaths.String() + "/" + fileName,
						Content: packageName,
					},
					Kind: helpers.KindOfFile(fileName),
				})
			}

			continue
		}

		if first {
			stackPaths.Push(line.info)
			stackIndents.Push(0)

			folder := folder.Folder{
				Path: stackPaths.String(),
			}
			items = append(items, item.Item{
				ObjectPath: folder,
				Kind:       helpers.KindOfFile(line.info),
			})

			first = false

			continue
		}

		if line.level > stackIndents.Peek().(int) {
			stackPaths.Push(line.info)
			stackIndents.Push(line.level)

			folder := folder.Folder{
				Path: stackPaths.String(),
			}
			items = append(items, item.Item{
				ObjectPath: folder,
				Kind:       helpers.KindOfFile(line.info),
			})

			continue
		}

		if line.level == stackIndents.Peek().(int) {
			stackPaths.Pop()

			stackPaths.Push(line.info)
			stackIndents.Push(line.level)

			items = append(items, item.Item{
				ObjectPath: folder.Folder{
					Path: stackPaths.String(),
				},
				Kind: helpers.KindOfFile(line.info),
			})

			continue
		}

		for line.level < stackIndents.Peek().(int) && len(stackIndents) > 1 {
			stackPaths.Pop()
			stackIndents.Pop()

			if line.level == stackIndents.Peek().(int) {
				stackPaths.Pop()
				stackIndents.Pop()

				break
			}
		}

		stackPaths.Push(line.info)
		stackIndents.Push(line.level)

		items = append(items, item.Item{
			ObjectPath: folder.Folder{
				Path: stackPaths.String(),
			},
			Kind: helpers.KindOfFile(line.info),
		})
	}

	return &items
}

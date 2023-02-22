package line

import (
	"fmt"
	"strings"

	"github.com/cezarovici/goLayouter/app/helpers"
	"github.com/cezarovici/goLayouter/app/stack"
	"github.com/cezarovici/goLayouter/domain"
	"github.com/cezarovici/goLayouter/domain/file"
	"github.com/cezarovici/goLayouter/domain/folder"
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

func (lines Lines) ToItems() *domain.Items {
	var items domain.Items

	first := true

	var stackPaths stack.Stack
	var stackIndents stack.Stack
	stackPackages := &stack.Stack{_defaultPackage}

	for _, line := range lines {
		switch helpers.TypeOfFile(line.info) {
		case "path":
			stackPackages = &stack.Stack{_defaultPackage}
			stackIndents = nil
			stackPaths = nil

			if !helpers.ToCurentDirectory(line.info) {
				stackPaths.Push(helpers.RemoveSelector(line.info))
				stackIndents.Push(-1)

				items = append(items, domain.Item{
					ObjectPath: &folder.Folder{
						Path: stackPaths.String(),
					},
					Kind: helpers.KindOfFile(line.info),
				})

				continue
			}

			stackIndents.Push(line.level)

			continue

		case "package":
			stackPackages.Push(helpers.RemoveSelector(line.info))

			continue

		case "file":
			packageName := helpers.GetRootPackage(stackPackages.String())

			if !stackPackages.IsEmpty() {
				packageName = stackPackages.Peek().(string)
			}

			if helpers.IsTestPackage(stackPackages.Peek().(string)) {
				testPackage := stackPackages.Peek()
				stackPackages.Pop()
				packageName = stackPackages.Peek().(string)

				stackPackages.Push(testPackage)
			}

			files := helpers.SplitLine(line.info, stackPackages.Peek().(string))
			for _, fileName := range files {
				if helpers.KindOfFile(fileName) == "main" {
					packageName = _defaultPackage
				}

				items = append(items, domain.Item{
					ObjectPath: file.File{
						Path:    stackPaths.String() + "/" + fileName,
						Content: packageName,
					},
					Kind: helpers.KindOfFile(fileName),
				})
			}

			stackIndents.Push(line.level)
			continue
		}

		if (stackIndents != nil) && stackIndents.Peek().(int) < 0 {
			items = items[:len(items)-1]
		}

		if first {
			stackPaths.Push(line.info)
			stackIndents.Push(line.level)

			folder := &folder.Folder{
				Path: stackPaths.String(),
			}
			items = append(items, domain.Item{
				ObjectPath: folder,
				Kind:       helpers.KindOfFile(folder.Path),
			})

			first = false

			continue
		}

		if line.level > stackIndents.Peek().(int) {
			stackPaths.Push(line.info)
			stackIndents.Push(line.level)

			folder := &folder.Folder{
				Path: stackPaths.String(),
			}
			items = append(items, domain.Item{
				ObjectPath: folder,
				Kind:       helpers.KindOfFile(folder.Path),
			})

			continue
		}

		if line.level == stackIndents.Peek().(int) {
			stackPaths.Pop()

			stackPaths.Push(line.info)
			stackIndents.Push(line.level)

			items = append(items, domain.Item{
				ObjectPath: folder.Folder{
					Path: stackPaths.String(),
				},
				Kind: helpers.KindOfFile(line.info),
			})

			continue
		}

		for line.level <= stackIndents.Peek().(int) && len(stackIndents) > 1 {
			stackPaths.Pop()
			stackIndents.Pop()
		}

		stackPaths.Push(line.info)
		stackIndents.Push(line.level)

		items = append(items, domain.Item{
			ObjectPath: folder.Folder{
				Path: stackPaths.String(),
			},
			Kind: helpers.KindOfFile(line.info),
		})
	}

	return &items
}

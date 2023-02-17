package line

import (
	"fmt"
	"strings"

	"github.com/cezarovici/goLayouter/app/helpers"
	"github.com/cezarovici/goLayouter/app/stack"
	"github.com/cezarovici/goLayouter/domain"
)

type Line struct {
	info  string
	level int
}

type Lines []Line

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

	var res Lines

	for _, line := range content {
		res = append(res, ConvertToLine(line))
	}

	return res, nil
}

func (lines Lines) ParseTo() []domain.Item {
	var items []domain.Item

	var stackPaths stack.Stack
	var stackIndents stack.Stack
	var stackPackages stack.Stack

	for index, line := range lines {
		switch helpers.TypeOfFile(line.info) {
		case "path":
			stackPackages = nil
			stackIndents = nil
			stackPaths = nil

			if !helpers.ToCurentDirectory(line.info) {
				stackPaths.Push(helpers.ReturnSelector(line.info))
				stackIndents.Push(-1)

				continue
			}

			stackIndents.Push(line.level)

			continue

		case "package":
			stackPackages.Push(helpers.ReturnSelector(line.info))

			continue

		case "file":

		}

	}

	return items
}

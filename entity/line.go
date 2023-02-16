package line

import (
	"fmt"
	"strings"
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

func (lines Lines) ParseTo() {
	//for index, line := range lines {

	//}
}

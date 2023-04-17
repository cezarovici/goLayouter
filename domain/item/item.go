package item

import (
	"strings"

	"github.com/cezarovici/goLayouter/domain"
)

type KindOfFile int8

const (
	Folder KindOfFile = iota
	NormalFile
	Main
	Test
	TableDriven
	Object
)

// Item represents a single item in a collection.
type Item struct {
	// ObjectPath represents the path to the file
	ObjectPath domain.FileOperations
	// Kind describes the type of the item.
	Kind KindOfFile
}

// Items represents a collection of items.
type Items []Item

// KindOfFile returns a string representing the kind of a file based on its name.
func NewKindOfFile(fileName string) KindOfFile {
	if !strings.Contains(fileName, ".") {
		return Folder
	}

	if fileName == "main" || fileName == "main.go" {
		return Main
	}

	if strings.Contains(fileName, "test") {
		return Test
	}

	if strings.Contains(fileName, "obj") {
		return Object
	}

	if strings.Contains(fileName, ".") {
		return NormalFile
	}

	return Folder
}

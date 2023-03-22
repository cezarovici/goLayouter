package item

import "github.com/cezarovici/goLayouter/domain"

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

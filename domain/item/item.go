package item

import "github.com/cezarovici/goLayouter/domain"

type Item struct {
	ObjectPath domain.FileOperations
	Kind       string
}

type Items []Item

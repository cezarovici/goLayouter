package item

import "github.com/cezarovici/goLayouter/domain"

type Item struct {
	ObjectPath domain.FileOperations
	Kind       string
}

type Items []Item

func (items *Items) ToStrings() []string {
	var res []string
	for _, item := range *items {
		res = append(res, item.ObjectPath.GetPath())
	}

	return res
}

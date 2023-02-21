package domain

type Item struct {
	ObjectPath FileOperations
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

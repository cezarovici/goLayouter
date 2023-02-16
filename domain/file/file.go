package file

import (
	"os"

	"github.com/cezarovici/goLayouter/domain"
)

type File struct {
	path    string
	content string
}

var _ domain.FileOperations = &File{}

func (f File) WriteToDisk() error {
	file, errCreate := os.Create(f.path)
	if errCreate != nil {
		return errCreate
	}

	_, errWrite := file.Write([]byte(f.content))
	if errWrite != nil {
		return errCreate
	}

	return nil
}

package file

import (
	"io"
	"os"

	"github.com/cezarovici/goLayouter/domain"
)

type File struct {
	Path    string
	Content string
}

var _ domain.FileOperations = &File{}
var _ io.Writer = &File{}

func (f File) Write([]byte) (int, error) {
	return 0, nil
}

func (f File) WriteToDisk() error {
	file, errCreate := os.Create(f.Path)
	if errCreate != nil {
		return errCreate
	}

	_, errWrite := file.Write([]byte(f.Content))
	if errWrite != nil {
		return errCreate
	}

	return nil
}

func (f File) GetPath() string {
	return f.Path
}

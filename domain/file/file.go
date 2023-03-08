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

func (f File) GetContent() []byte {
	return []byte(f.Content)
}

func (f File) Write([]byte) (int, error) {
	file, errCreate := os.Create(f.Path)
	if errCreate != nil {
		return 0, errCreate
	}

	length, errWrite := file.Write([]byte(f.Content))
	if errWrite != nil {
		return 0, errWrite
	}

	return length, nil
}

func (f File) GetPath() string {
	return f.Path
}

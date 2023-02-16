package file

import (
	"os"
)

type File struct {
	path    string
	content string
}

var _ interfaces.FileOperations = &File{}

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

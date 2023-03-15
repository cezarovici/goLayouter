package folder

import (
	"os"

	"github.com/cezarovici/goLayouter/domain"
)

type Folder struct {
	Path string
}

var _ domain.FileOperations = &Folder{}

func (f Folder) GetPackage() []byte {
	return nil
}

func (f Folder) Write([]byte) error {
	err := os.MkdirAll(f.Path, 0755)
	if err != nil {
		return err
	}
	return nil
}

func (f Folder) GetPath() string {
	return f.Path
}

package folder

import (
	"os"

	"github.com/cezarovici/goLayouter/domain"
)

type Folder struct {
	Path string
}

var _ domain.FileOperations = &Folder{}

func (f Folder) WriteToDisk() error {
	return os.MkdirAll(f.Path, 0755)
}

func (f Folder) GetPath() string {
	return f.Path
}

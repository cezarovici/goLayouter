package folder

import (
	"os"

	"github.com/cezarovici/goLayouter/domain"
)

type Folder struct {
	path string
}

var _ domain.FileOperations = &Folder{}

func (f Folder) WriteToDisk() error {
	return os.MkdirAll(f.path, 0755)
}

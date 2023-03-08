package folder

import (
	"io"
	"os"

	"github.com/cezarovici/goLayouter/domain"
)

type Folder struct {
	Path string
}

var _ io.Writer = &Folder{}

var _ domain.FileOperations = &Folder{}

func (f Folder) Write([]byte) (int, error) {
	return 0, nil
}
func (f Folder) WriteToDisk() error {
	return os.MkdirAll(f.Path, 0755)
}

func (f Folder) GetPath() string {
	return f.Path
}

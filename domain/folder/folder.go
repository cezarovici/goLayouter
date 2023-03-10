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

func (f Folder) GetPackage() []byte {
	return nil
}

func (f Folder) Write([]byte) (int, error) {
	err := os.MkdirAll(f.Path, 0755)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

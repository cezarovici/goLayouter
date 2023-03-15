package folder

import (
	"os"

	apperrors "github.com/cezarovici/goLayouter/app/errors"
	"github.com/cezarovici/goLayouter/domain"
)

type Folder struct {
	Path string
}

var _ domain.FileOperations = &Folder{}

func (f Folder) GetContent() []byte {
	return []uint8([]byte(nil))
}

func (f Folder) Write([]byte) error {
	errWrite := os.MkdirAll(f.Path, 0755)
	if errWrite != nil {
		return &apperrors.ErrDomain{
			Caller:     "object file -> Write",
			MethodName: "file.Writes",
			Issue:      errWrite,
		}
	}

	return nil
}

func (f Folder) GetPath() string {
	return f.Path
}

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
	return ([]byte(nil))
}

func (f Folder) Write([]byte) error {
	errWrite := os.MkdirAll(f.Path, 0o755)
	if errWrite != nil {
		return &apperrors.DomainError{
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

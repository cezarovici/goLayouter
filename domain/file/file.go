package file

import (
	"os"

	apperrors "github.com/cezarovici/goLayouter/app/errors"
	"github.com/cezarovici/goLayouter/domain"
)

// File represents a file on disk and implements the domain.FileOperations interface.
type File struct {
	Path       string // Path to the file on disk.
	Package    string // The Package of the file.
	ObjectName string
}

// Ensure File implements the domain.FileOperations interface.
var _ domain.FileOperations = &File{}

// GetContent returns the Package of the file as a byte slice.
func (f File) GetContent() []byte {
	return []byte(f.Package)
}

// Write writes the Package of the file to disk at the specified path.
func (f File) Write(_ []byte) error {
	// Create the file at the specified path.
	file, errCreate := os.Create(f.Path)
	if errCreate != nil {
		return &apperrors.DomainError{
			Caller:     "object file -> Write",
			MethodName: "os.Create",
			Issue:      errCreate,
		}
	}

	// Write the Package of the file to disk.
	_, errWrite := file.Write([]byte(f.Package))
	if errWrite != nil {
		return &apperrors.DomainError{
			Caller:     "object file -> Write",
			MethodName: "file.Writes",
			Issue:      errWrite,
		}
	}

	// Return no error
	return nil
}

func (f File) GetPath() string {
	return f.Path
}

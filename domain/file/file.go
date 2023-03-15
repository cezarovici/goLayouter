package file

import (
	"os"

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

// GetPackage returns the Package of the file as a byte slice.
func (f File) GetPackage() []byte {
	return []byte(f.Package)
}

// Write writes the Package of the file to disk at the specified path.
func (f File) Write(Package []byte) error {
	// Create the file at the specified path.
	file, errCreate := os.Create(f.Path)
	if errCreate != nil {
		return errCreate
	}

	// Write the Package of the file to disk.
	_, errWrite := file.Write([]byte(f.Package))
	if errWrite != nil {
		return errWrite
	}

	// Return the number of bytes written.
	return nil
}

func (f File) GetPath() string {
	return f.Path
}

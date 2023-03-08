package file

import (
	"io"
	"os"

	"github.com/cezarovici/goLayouter/domain"
)

// File represents a file on disk and implements the domain.FileOperations interface.
type File struct {
	Path    string // Path to the file on disk.
	Content string // The content of the file.
}

// Ensure File implements the domain.FileOperations interface.
var _ domain.FileOperations = &File{}

// Ensure File implements the io.Writer interface.
var _ io.Writer = &File{}

// GetContent returns the content of the file as a byte slice.
func (f File) GetContent() []byte {
	return []byte(f.Content)
}

// Write writes the content of the file to disk at the specified path.
func (f File) Write(content []byte) (int, error) {
	// Create the file at the specified path.
	file, errCreate := os.Create(f.Path)
	if errCreate != nil {
		return 0, errCreate
	}

	// Write the content of the file to disk.
	length, errWrite := file.Write([]byte(f.Content))
	if errWrite != nil {
		return 0, errWrite
	}

	// Return the number of bytes written.
	return length, nil
}

// GetPath returns the path of the file.
func (f File) GetPath() string {
	return f.Path
}

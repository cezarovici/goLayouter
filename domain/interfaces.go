package domain

// FileOperations is an interface that defines file operations.
type FileOperations interface {
	// GetPackage returns the content of a file as a byte slice.
	GetPackage() []byte
	// Write writes the given byte slice to a file and returns the number of bytes written and an error (if any).
	Write([]byte) (int, error)
}

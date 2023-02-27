package domain

type FileOperations interface {
	WriteToDisk() error
	GetPath() string
}

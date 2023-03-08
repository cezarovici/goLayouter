package domain

type FileOperations interface {
	GetContent() []byte
	Write([]byte) (int, error)
}

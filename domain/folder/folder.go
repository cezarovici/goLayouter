package folder

import "os"

type Folder struct {
	path string
}

var _ interfaces.FileOperations = &Folder{}

func (f Folder) WriteToDisk() error {
	return os.MkdirAll(f.path, 0755)
}

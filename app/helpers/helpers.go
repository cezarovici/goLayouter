package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Read file from parsed path
func ReadFile(filePath string) ([]string, error) {
	_, errExists := os.Stat(filePath)
	if errExists != nil {
		return nil, fmt.Errorf("not a valid file parsed")

	}

	fileHandler, errOp := os.Open(filePath)
	if errOp != nil {
		return nil, errOp
	}

	var errClo error
	defer func() {
		errClo = fileHandler.Close()
	}()

	var res []string

	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("empty file passed")
	}

	return res, errClo
}

// Return the type of file
func TypeOfFile(fileName string) string {
	switch {
	case strings.Contains(fileName, "!"):
		return "path"
	case strings.Contains(fileName, "."):
		return "file"
	case strings.Contains(fileName, "#"):
		return "package"
	default:
		return "folder"
	}
}

package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadFile reads a file from the parsed file path
// and returns its contents as a slice of strings.
func ReadFile(filePath string) ([]string, error) {
	// Check if the file exists
	if _, errExists := os.Stat(filePath); errExists != nil {
		return nil, fmt.Errorf("not a valid file parsed")
	}

	// Open the file
	fileHandler, errOp := os.Open(filePath)
	if errOp != nil {
		return nil, errOp
	}

	// Close the file when the function exits
	defer fileHandler.Close()

	var res []string

	// Read the file line by line and append each line to the result slice
	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	// If the file is empty, return an error
	if len(res) == 0 {
		return nil, fmt.Errorf("empty file passed")
	}

	return res, nil
}

// TypeOfFile returns the type of a file based on its name.
func TypeOf(fileName string) string {
	switch {
	case strings.Contains(fileName, "!"):
		return "path"
	case strings.Contains(fileName, "."):
		return "file"
	case strings.Contains(fileName, "#"):
		return "package"
	case len(fileName) == 0:
		return "empty"
	default:
		return "folder"
	}
}

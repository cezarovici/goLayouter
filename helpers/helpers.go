package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

// ReadFile reads a file from the parsed file path and returns its contents as a slice of strings.
func ReadFile(filePath string) ([]string, error) {
	// Check if the file exists
	_, errExists := os.Stat(filePath)
	if errExists != nil {
		return nil, fmt.Errorf("not a valid file parsed")
	}

	// Open the file
	fileHandler, errOp := os.Open(filePath)
	if errOp != nil {
		return nil, errOp
	}

	// Close the file when the function exits
	var errClo error
	defer func() {
		errClo = fileHandler.Close()
	}()

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

	return res, errClo
}

// TypeOfFile returns the type of a file based on its name.
func TypeOfFile(fileName string) string {
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

// ToCurentDirectory returns true if the given path moves to the current directory.
func ToCurentDirectory(pathName string) bool {
	return RemoveSelector(pathName) == "."
}

// RemoveSelector removes the leading selector from a line of text.
func RemoveSelector(line string) string {
	return line[2:]
}

// KindOfFile returns a string representing the kind of a file based on its name.
func KindOfFile(fileName string) string {
	if fileName == "main" || fileName == "main.go" {
		return "main"
	}

	if strings.Contains(fileName, "obj") {
		return "object"
	}

	if strings.Contains(fileName, "test") {
		return "test"
	}

	if strings.Contains(fileName, ".") {
		return "normalFile"
	}

	return "folder"
}

// IsTestPackage returns true if the given package name indicates a test package.
func IsTestPackage(packageName string) bool {
	return packageName == "t" || packageName == "tt"
}

// CreateGolangTestFile returns the file path of the test file corresponding to a given file path.
func CreateGolangTestFile(text string) (string, error) {
	path, fileName := path.Split(text)

	pos := strings.Index(fileName, ".")
	if pos == -1 {
		return "", errors.New("passed is not a file path")
	}

	return path + fileName[:pos] + "_test.go", nil
}

// SplitLine splits a line of text into a slice of file paths, and also creates test files for test packages.
func SplitLine(text, packageName string) []string {
	var res []string
	files := strings.Split(text, " ")

	for _, file := range files {
		fileTrimmed := strings.TrimLeft(file, " ")
		if len(fileTrimmed) == 0 {
			continue
		}

		if IsTestPackage(packageName) {
			testFile, err := CreateGolangTestFile(fileTrimmed)

			if err == nil {
				res = append(res, fileTrimmed, testFile)
			}

			continue
		}

		res = append(res, fileTrimmed)
	}

	return res
}

// GetLastPath returns the package name of the last directory in the given path.
// If the path is empty, it returns "package main"
func GetLastPath(path string) string {
	if len(path) == 0 {
		return "package main"
	}

	// Split the input path by the "/" separator
	folders := strings.Split(path, "/")

	// Get the last folder in the path
	lastFolder := folders[len(folders)-1]

	// Return the package declaration string for the last folder
	return fmt.Sprintf("package %s", lastFolder)

}

func RemoveObjectKey(fileName string) string {
	if fileName == "obj.go" || fileName == "object.go" || !strings.Contains(fileName, "_") {
		return fileName
	}

	newFileName, case1 := strings.CutPrefix(fileName, "object_")
	if !case1 {
		newFileName, _ = strings.CutPrefix(fileName, "obj_")
	}

	return newFileName
}

package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
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

// Return a true if the path move you to the curent directory
func ToCurentDirectory(pathName string) bool {
	return pathName[2:] == "."
	//! . -> true
	//! another directory -> false
}

func RemoveSelector(line string) string {
	return line[2:]
}

func KindOfFile(fileName string) string {
	if strings.Contains(fileName, "main") {
		return "main"
	}

	if strings.Contains(fileName, "test") {
		return "test"
	}

	if strings.Contains(fileName, "obj") {
		return "object"
	}

	if strings.Contains(fileName, ".") {
		return "normalFile"
	}

	return "folder"
}

func IsTestPackage(packageName string) bool {
	packageName = RemoveSelector(packageName)

	return packageName == "t" || packageName == "tt"
}

func CreateGolangTestFile(text string) (string, error) {
	path, fileName := path.Split(text)

	pos := strings.Index(fileName, ".")
	if pos == -1 {
		return "", errors.New("passed is not a file path")
	}

	return path + fileName[:pos] + "_test.go", nil
}

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

func GetRootPackage(pathName string) string {
	if len(pathName) == 0 {
		return "package main"
	}
	_, fileName := path.Split(pathName)

	return fileName
}

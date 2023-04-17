package helpers

import (
	"errors"
	"fmt"
	"path"
	"strings"
)

// SplitLine splits a line of text into a slice of file paths,
// and also creates test files for test packages.
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
			if err != nil {
				return nil
			}

			res = append(res, fileTrimmed, testFile)

			continue
		}

		res = append(res, fileTrimmed)
	}

	return res
}

// GetPackageFromPath returns the package name of the last directory in the given path.
// If the path is empty, it returns "package main".
func GetPackageFrom(path string) string {
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

func RemoveObjectPrefix(fileName string) string {
	isObject := strings.Contains(fileName, "obj")
	if !isObject {
		return fileName
	}

	if fileName == "obj.go" || fileName == "object.go" {
		return fileName
	}

	newFileName, case1 := strings.CutPrefix(fileName, "object")
	if !case1 {
		newFileName, _ = strings.CutPrefix(fileName, "obj")
	}

	fileName = strings.Replace(newFileName, "_", "", 1)

	return fileName
}

// ToCurentDirectory returns true if the given path moves to the current directory.
func ToCurentDirectory(pathName string) bool {
	return RemoveSelector(pathName) == "."
}

// RemoveSelector removes the leading selector from a line of text.
func RemoveSelector(line string) string {
	return line[2:]
}

// IsTestPackage returns true if the given package name indicates a test package.
func IsTestPackage(packageName string) bool {
	return packageName == "t" || packageName == "tt"
}

// CreateGolangTestFile returns the file path of the test
// file corresponding to a given file path.
func CreateGolangTestFile(text string) (string, error) {
	path, fileName := path.Split(text)

	pos := strings.Index(fileName, ".")
	if pos == -1 {
		return "", errors.New("passed is not a file path")
	}

	return path + fileName[:pos] + "_test.go", nil
}

package services

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/cezarovici/goLayouter/domain/file"
	"github.com/cezarovici/goLayouter/domain/folder"
	"github.com/cezarovici/goLayouter/domain/item"
	"github.com/stretchr/testify/require"
)

func TestRender(t *testing.T) {
	// Define some test data
	filePath := "test.txt"
	fileContent := []byte("hello world")
	folderPath := "testFolder"

	items := item.Items{
		item.Item{
			Kind:       "file",
			ObjectPath: file.File{Path: filePath, Content: string(fileContent)},
		},
		item.Item{
			Kind:       "folder",
			ObjectPath: folder.Folder{Path: folderPath},
		},
	}

	renderFuncs := map[string]func(io.Writer, any) error{
		"file":   func(w io.Writer, data any) error { return nil },
		"folder": func(w io.Writer, data any) error { return nil },
	}

	// Create a new service instance
	serv := Service{
		paths:       items,
		renderFuncs: renderFuncs,
	}

	// Call the Render method
	err := serv.Render()

	// Check that there were no errors
	require.NoError(t, err)

	// Check that the file was written correctly
	outputContent, errRead := ioutil.ReadFile(filePath)
	require.NoError(t, errRead)
	require.Equal(t, fileContent, outputContent)

	// Check that the folder was created
	_, errStat := os.Stat(folderPath)
	require.NoError(t, errStat)

	// Clean up the test data
	require.NoError(t, os.Remove(filePath))
	require.NoError(t, os.Remove(folderPath))
}

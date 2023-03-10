package services

import (
	"os"
	"testing"

	"github.com/cezarovici/goLayouter/app/services/renders"
	"github.com/cezarovici/goLayouter/domain/file"
	"github.com/cezarovici/goLayouter/domain/folder"
	"github.com/cezarovici/goLayouter/domain/item"
	"github.com/stretchr/testify/require"
)

func TestRender(t *testing.T) {
	// Define some test data
	filePath := "main.go"
	folderPath := "testFolder"

	items := item.Items{
		item.Item{
			Kind:       "main",
			ObjectPath: file.File{Path: filePath}, //Package: string(filePackage)},
		},
		item.Item{
			Kind:       "folder",
			ObjectPath: folder.Folder{Path: folderPath},
		},
	}

	// Create a new service instance
	serv, errNewService := NewService(items, renders.RenderFuncs)
	require.NoError(t, errNewService)

	// Call the Render method
	err := serv.Render()

	// Check that there were no errors
	require.NoError(t, err)

	// Check that the folder was created
	_, errStat := os.Stat(folderPath)
	require.NoError(t, errStat)

	// Clean up the test data
	require.NoError(t, os.Remove(filePath))
	require.NoError(t, os.Remove(folderPath))
}

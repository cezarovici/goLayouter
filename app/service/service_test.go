package service_test

import (
	"errors"
	"os"
	"testing"

	apperrors "github.com/cezarovici/goLayouter/app/errors"
	"github.com/cezarovici/goLayouter/app/service"
	"github.com/cezarovici/goLayouter/app/service/render"
	"github.com/cezarovici/goLayouter/domain/file"
	"github.com/cezarovici/goLayouter/domain/folder"
	"github.com/cezarovici/goLayouter/domain/item"
	"github.com/stretchr/testify/require"
)

func TestRender(t *testing.T) {
	// Define some test data
	const (
		filePath    = "main.go"
		folderPath  = "testFolder"
		filePackage = "package file"
	)

	items := item.Items{
		item.Item{
			Kind:       item.Main,
			ObjectPath: file.File{Path: filePath, Package: string(filePackage)},
		},
		item.Item{
			Kind:       item.Folder,
			ObjectPath: folder.Folder{Path: folderPath},
		},
	}

	// Create a new service instance
	serv, errNewService := service.NewService(items, render.Funcs)
	require.NoError(t, errNewService)

	// Call the Render method
	// Check that there were no errors
	require.NoError(t, serv.Render())

	// Check that the folder was created
	_, errStat := os.Stat(folderPath)
	require.NoError(t, errStat)

	// Clean up the test data
	require.NoError(t, os.Remove(filePath))
	require.NoError(t, os.Remove(folderPath))
}

func TestNewService(t *testing.T) {
	items := item.Items{
		{ObjectPath: file.File{Path: "/path/to/template1.tmpl", Package: "Template 1 data"}, Kind: item.NormalFile},
		{ObjectPath: file.File{Path: "/path/to/template2.tmpl", Package: "Template 2 data"}, Kind: item.NormalFile},
	}
	// Test with non-empty items and renders map
	_, err := service.NewService(items, render.Funcs)
	require.NoError(t, err)

	// Test with empty items
	_, err = service.NewService(item.Items{}, render.Funcs)
	require.Error(t, err)

	expectedErr := &apperrors.ServiceError{
		Caller:     "Service",
		MethodName: "NewService",
		Issue:      errors.New("no items parsed"),
	}
	if !errors.As(err, &expectedErr) {
		t.Errorf("NewService() error = %v, expected %v", err, expectedErr)
	}
}

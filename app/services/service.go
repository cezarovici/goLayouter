package services

import (
	"errors"
	"log"

	"github.com/cezarovici/goLayouter/domain/item"
)

// Service represents a service that renders items to the filesystem.
type Service struct {
	paths       item.Items
	renderFuncs map[string]func(string, any) error
}

// NewService creates a new Service instance.
func NewService(items item.Items, renders map[string]func(string, any) error) (*Service, error) {
	if len(items) == 0 {
		return nil, errors.New("no items provided")
	}

	return &Service{
		paths:       items,
		renderFuncs: renders,
	}, nil
}

// RenderItems renders all items to the filesystem.
func (service Service) Render() error {
	for _, path := range service.paths {
		if errWrite := path.ObjectPath.Write(path.ObjectPath.GetPackage()); errWrite != nil {
			return errWrite
		}

		if path.Kind == "normalFile" || path.Kind == "folder" {
			continue
		}
		log.Print(path.ObjectPath)
		//TODO no object name recognised
		errRender := service.renderFuncs[path.Kind](path.ObjectPath.GetPath(), path.ObjectPath)
		if errRender != nil {
			return errRender
		}
	}

	return nil
}

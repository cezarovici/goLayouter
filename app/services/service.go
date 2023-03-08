package services

import (
	"errors"
	"io"

	"github.com/cezarovici/goLayouter/app/services/renders"
	"github.com/cezarovici/goLayouter/domain/item"
)

// Service represents a service that renders items to the filesystem.
type Service struct {
	paths       item.Items
	renderFuncs map[string]func(io.Writer, any) error
}

// NewService creates a new Service instance.
func NewService(items item.Items) (*Service, error) {
	if len(items) == 0 {
		return nil, errors.New("no items provided")
	}

	return &Service{
		paths:       items,
		renderFuncs: renders.RenderFuncs,
	}, nil
}

// RenderItems renders all items to the filesystem.
func (service Service) Render() error {
	for _, path := range service.paths {
		_, errWrite := path.ObjectPath.Write(path.ObjectPath.GetContent())
		if errWrite != nil {
			return errWrite
		}

		if path.Kind != "folder" {
			service.renderFuncs[path.Kind](path.ObjectPath, nil)
			//TODO: replace nil with an extra object
		}
	}

	return nil
}

package services

import (
	"errors"

	apperrors "github.com/cezarovici/goLayouter/app/errors"
	"github.com/cezarovici/goLayouter/app/services/renders"
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
		return nil, &apperrors.ErrService{
			Caller:     "Service",
			MethodName: "NewService",
			Issue:      errors.New("no items parsed"),
			WithTime:   true,
		}
	}

	return &Service{
		paths:       items,
		renderFuncs: renders,
	}, nil
}

// RenderItems renders all items to the filesystem.
func (service Service) Render() error {
	for _, path := range service.paths {
		if errWrite := path.ObjectPath.Write(path.ObjectPath.GetContent()); errWrite != nil {
			return &apperrors.ErrService{
				Caller:     "Service",
				MethodName: "Render -> Write",
				Issue:      errWrite,
				WithTime:   false,
			}
		}

		if path.Kind == "normalFile" || path.Kind == "folder" {
			continue
		}

		errRender := renders.RenderFuncs[path.Kind](path.ObjectPath.GetPath(), path.ObjectPath)
		if errRender != nil {
			return &apperrors.ErrService{
				Caller:     "Service",
				MethodName: "Render -> render funcs",
				Issue:      errRender,
				WithTime:   true,
			}
		}
	}

	return nil
}

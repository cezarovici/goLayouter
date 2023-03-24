package service

import (
	"errors"

	apperrors "github.com/cezarovici/goLayouter/app/errors"
	"github.com/cezarovici/goLayouter/app/service/render"
	"github.com/cezarovici/goLayouter/domain/item"
)

// Service represents a service that renders items to the filesystem.
type Service struct {
	paths       item.Items
	renderFuncs map[item.KindOfFile]func(string, any) error
}

// NewService creates a new Service instance.
func NewService(items item.Items, renders map[item.KindOfFile]func(string, any) error) (*Service, error) {
	if len(items) == 0 {
		return nil, &apperrors.ServiceError{
			Caller:     "Service",
			MethodName: "NewService",
			Issue:      errors.New("no items parsed"),
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
			return &apperrors.ServiceError{
				Caller:     "Service",
				MethodName: "Render -> Write",
				Issue:      errWrite,
			}
		}

		if path.Kind == item.NormalFile || path.Kind == item.Folder {
			continue
		}

		errRender := render.Funcs[path.Kind](path.ObjectPath.GetPath(), path.ObjectPath)
		if errRender != nil {
			return &apperrors.ServiceError{
				Caller:     "Service",
				MethodName: "Render -> render funcs",
				Issue:      errRender,
			}
		}
	}

	return nil
}

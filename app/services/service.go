package services

import (
	"errors"
	"io"

	"github.com/cezarovici/goLayouter/domain/item"
)

type Service struct {
	paths       item.Items
	renderFuncs map[string]func(io.Writer, any) error
}

func NewService(content item.Items) (*Service, error) {
	if len(content) == 0 {
		return nil, errors.New("parsed content is empty")
	}

	return &Service{
		paths:       content,
		renderFuncs: _renderFuncs,
	}, nil
}

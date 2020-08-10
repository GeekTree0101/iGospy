package worker

import (
	"fmt"

	"github.com/GeekTree0101/iSpygo/model"
)

// BuilderError is error for builder
type BuilderError struct {
	desc string
}

// Error is builder error constructor
func (be *BuilderError) Error() string {
	return fmt.Sprintf("%s", be.desc)
}

// Builder is spy object generator
type Builder struct {
	node model.Node
}

// NewBuilder is builder constructor
func NewBuilder(node model.Node) Builder {
	return Builder{
		node: node,
	}
}

// GetPresenter return presenter spy object
func (b *Builder) GetPresenter() (string, error) {

	return "", &BuilderError{
		desc: "unknown error",
	}
}

// GetDisplayer return displayer spy object
func (b *Builder) GetDisplayer() (string, error) {

	return "", &BuilderError{
		desc: "unknown error",
	}
}

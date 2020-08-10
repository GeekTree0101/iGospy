package worker

import (
	"fmt"
	"strings"

	"github.com/GeekTree0101/iSpygo/model"
)

// ObjectType is CleanSwift object component type
type ObjectType int

const (
	// Interactor is business object
	Interactor = 0
	// Presenter is presentation object
	Presenter = 1
	// Displayer is display object
	Displayer = 2
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
	return b.getSpyObject(Presenter)
}

// GetDisplayer return displayer spy object
func (b *Builder) GetDisplayer() (string, error) {
	return b.getSpyObject(Displayer)
}

func (b *Builder) getSpyObject(objType ObjectType) (string, error) {

	var out string

	title := b.node.Name

	for _, us := range b.node.Children {

		if us.Type != model.Context {
			return "", &BuilderError{
				desc: "invalid structure",
			}
		}

		usecase := us.Name

		if len(us.Children) < 3 {
			return "", &BuilderError{
				desc: "invalid structure",
			}
		}

		behaivor := us.Children[objType].Name

		temp, err := b.getTemplate(objType)

		if err != nil {
			return "", err
		}

		temp = strings.ReplaceAll(temp, "<title>", title)
		temp = strings.ReplaceAll(temp, "<usecase>", usecase)
		temp = strings.ReplaceAll(temp, "<behavior>", behaivor)

		out += temp
	}

	if len(out) == 0 {
		return "", &BuilderError{
			desc: "output is empty",
		}
	}

	return out, nil
}

func (b *Builder) getTemplate(objType ObjectType) (string, error) {
	switch objType {
	case Presenter:
		return `var present<usecase>Called: Int = 0
	var present<usecase><behavior>: <title>.<usecase>.<behavior>?
	func present<usecase>(res: <title>.<usecase>.<behavior>) {
	  self.present<usecase>Called += 1
	  self.present<usecase><behavior> = res
	}
		
	`, nil

	case Displayer:
		return `var display<usecase><behavior>: <title>.<usecase>.<behavior>?
  func display<usecase>(viewModel: <title>.<usecase>.<behavior>) {
    self.display<usecase><behavior> = viewModel
  }

`, nil

	default:
		return "", &BuilderError{
			desc: "Unsupported object type",
		}
	}
}

package worker

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/GeekTree0101/iGospy/model"
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
	usecase model.Usecase
}

// NewBuilder is builder constructor
func NewBuilder(usecase model.Usecase) Builder {
	return Builder{
		usecase: usecase,
	}
}

// GetInteractor return interactor spy object
func (b *Builder) GetInteractor() (string, error) {
	return b.getSpyObject(Interactor)
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

	title := b.usecase.Title

	for _, ctx := range b.usecase.Contexts {

		temp, err := b.getTemplate(objType)

		if err != nil {
			return "", err
		}

		temp = strings.ReplaceAll(temp, "<title>", title)
		temp = strings.ReplaceAll(temp, "<usecase>", ctx)
		temp = strings.ReplaceAll(temp, "<lower_capitalize_first_letter_usecase>", b.makeLowerCapitalizeFirstLetter(ctx))

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
	case Interactor:
		return `var <lower_capitalize_first_letter_usecase>Called: Int = 0
var <lower_capitalize_first_letter_usecase>Req: <title>.<usecase>.Request?
func <lower_capitalize_first_letter_usecase>(req: <title>.<usecase>.Request) {
  self.<lower_capitalize_first_letter_usecase>Called += 1
  self.<lower_capitalize_first_letter_usecase>Req = req
}

`, nil
	case Presenter:
		return `var present<usecase>Called: Int = 0
var present<usecase>Response: <title>.<usecase>.Response?
func present<usecase>(res: <title>.<usecase>.Response) {
  self.present<usecase>Called += 1
  self.present<usecase>Response = res
}

`, nil

	case Displayer:
		return `var display<usecase>ViewModel: <title>.<usecase>.ViewModel?
func display<usecase>(viewModel: <title>.<usecase>.ViewModel) {
  self.display<usecase>ViewModel = viewModel
}

`, nil

	default:
		return "", &BuilderError{
			desc: "Unsupported object type",
		}
	}
}

func (b *Builder) makeLowerCapitalizeFirstLetter(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return str
}

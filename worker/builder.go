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
func (b *Builder) GetInteractor() (model.BuilderResult, error) {

	ig, err := b.getInterface(Interactor)

	if err != nil {
		return model.BuilderResult{}, err
	}

	sg, err := b.getSpyObject(Interactor)

	if err != nil {
		return model.BuilderResult{}, err
	}

	return model.BuilderResult{
		InterfaceGroup: ig,
		SpyGroup:       sg,
	}, nil
}

// GetPresenter return presenter spy object
func (b *Builder) GetPresenter() (model.BuilderResult, error) {

	ig, err := b.getInterface(Presenter)

	if err != nil {
		return model.BuilderResult{}, err
	}

	sg, err := b.getSpyObject(Presenter)

	if err != nil {
		return model.BuilderResult{}, err
	}

	return model.BuilderResult{
		InterfaceGroup: ig,
		SpyGroup:       sg,
	}, nil
}

// GetDisplayer return displayer spy object
func (b *Builder) GetDisplayer() (model.BuilderResult, error) {

	ig, err := b.getInterface(Displayer)

	if err != nil {
		return model.BuilderResult{}, err
	}

	sg, err := b.getSpyObject(Displayer)

	if err != nil {
		return model.BuilderResult{}, err
	}

	return model.BuilderResult{
		InterfaceGroup: ig,
		SpyGroup:       sg,
	}, nil
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
var <lower_capitalize_first_letter_usecase>Request: <title>.<usecase>.Request?

func <lower_capitalize_first_letter_usecase>(request: <title>.<usecase>.Request) {
  self.<lower_capitalize_first_letter_usecase>Called += 1
  self.<lower_capitalize_first_letter_usecase>Request = request
}

`, nil
	case Presenter:
		return `var present<usecase>Called: Int = 0
var present<usecase>Response: <title>.<usecase>.Response?

func present<usecase>(response: <title>.<usecase>.Response) {
  self.present<usecase>Called += 1
  self.present<usecase>Response = response
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

func (b *Builder) getInterface(objType ObjectType) (string, error) {

	var out string

	title := b.usecase.Title

	for _, ctx := range b.usecase.Contexts {

		temp, err := b.getInterfaceTemplate(objType)

		if err != nil {
			return "", err
		}

		temp = strings.ReplaceAll(temp, "<title>", title)
		temp = strings.ReplaceAll(temp, "<usecase>", ctx)
		temp = strings.ReplaceAll(temp, "<lower_capitalize_first_letter_usecase>", b.makeLowerCapitalizeFirstLetter(ctx))

		out += temp + "\n"
	}

	if len(out) == 0 {
		return "", &BuilderError{
			desc: "output is empty",
		}
	}

	return out, nil
}

func (b *Builder) getInterfaceTemplate(objType ObjectType) (string, error) {
	switch objType {
	case Interactor:
		return `func <lower_capitalize_first_letter_usecase>(request: <title>.<usecase>.Request)`, nil
	case Presenter:
		return `func present<usecase>(response: <title>.<usecase>.Response)`, nil
	case Displayer:
		return `func display<usecase>(viewModel: <title>.<usecase>.ViewModel)`, nil

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

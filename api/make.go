package api

import (
	"net/http"

	"github.com/GeekTree0101/iGospy/model"

	"github.com/GeekTree0101/iGospy/worker"

	"github.com/labstack/echo/v4"
)

// Make is spy object generator API
type Make struct {
}

// NewMake is constructor
func NewMake() Make {
	return Make{}
}

// GetIndexPage return index.html
func (m *Make) GetIndexPage(c echo.Context) error {

	return c.Render(http.StatusOK, "index.html", nil)
}

// PostMake return generated spy object
func (m *Make) PostMake(c echo.Context) error {

	req := new(model.MakeRequest)

	err := c.Bind(req)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	parser := worker.NewParser()
	usecase, err := parser.MakeUsecase(req.Usecase)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	builder := worker.NewBuilder(usecase)

	interactor, err := builder.GetInteractor()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	presenter, err := builder.GetPresenter()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	displayer, err := builder.GetDisplayer()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	res := model.MakeResponse{
		Interactor:          interactor.SpyGroup,
		InteractorInterface: interactor.InterfaceGroup,
		Presenter:           presenter.SpyGroup,
		PresenterInterface:  presenter.InterfaceGroup,
		Displayer:           displayer.SpyGroup,
		DisplayerInterface:  displayer.InterfaceGroup,
	}

	return c.JSON(http.StatusOK, res)
}

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
	usecase := c.FormValue("usecase")
	parser := worker.NewParser()
	nodes, err := parser.Processing(usecase)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	builder := worker.NewBuilder(nodes)

	presenter, err := builder.GetPresenter()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	displayer, err := builder.GetDisplayer()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	res := model.MakeResponse{
		Presenter: presenter,
		Displayer: displayer,
	}

	return c.JSON(http.StatusOK, res)
}

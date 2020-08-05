package server

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"text/template"

	"github.com/GeekTree0101/iSpygo/util"
	"github.com/labstack/echo/v4"
)

// Server is server object
type Server struct {
	serv *echo.Echo
	port string
}

// NewServer is server constructor
func NewServer(port string) Server {
	return Server{
		serv: echo.New(),
		port: port,
	}
}

// Mount is configurate some modules for server
func (s *Server) Mount() {
	renderer := &util.Template{
		Templates: template.Must(template.ParseGlob("static/public/*.html")),
	}
	s.serv.Renderer = renderer
}

// Route is server route configuration
func (s *Server) Route() {

	s.serv.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})
}

// Run is execute server
func (s *Server) Run() {

	s.serv.Logger.Fatal(s.serv.Start(":" + s.port))

	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", "https://localhost:"+s.port).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://localhost:"+s.port).Start()
	case "darwin":
		err = exec.Command("open", "https://localhost:"+s.port).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Fatal(err)
	}
}

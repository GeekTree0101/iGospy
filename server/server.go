package server

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"text/template"

	"github.com/GeekTree0101/iSpygo/api"
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
		Templates: template.Must(template.ParseGlob("app/public/*.html")),
	}
	s.serv.Renderer = renderer

	s.serv.File("/global.css", "app/public/global.css")
	s.serv.File("/build/bundle.css", "app/public/build/bundle.css")
	s.serv.File("/build/bundle.js", "app/public/build/bundle.js")
}

// Route is server route configuration
func (s *Server) Route() {

	makeAPI := api.NewMake()

	s.serv.GET("/", makeAPI.GetIndexPage)
	s.serv.POST("/make", makeAPI.PostMake)
}

// Run is execute server
func (s *Server) Run() {

	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", "http://localhost:"+s.port).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:"+s.port).Start()
	case "darwin":
		err = exec.Command("open", "http://localhost:"+s.port).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Fatal(err)
	}

	s.serv.Logger.Fatal(s.serv.Start(":" + s.port))
}

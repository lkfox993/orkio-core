package http

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	server := &Server{
		echo: e,
	}

	return server
}

func (s *Server) Start(addr string) error {
	s.registerRoutes()

	return s.echo.Start(addr)
}

func (s *Server) Shutdown() error {
	return s.echo.Close()
}

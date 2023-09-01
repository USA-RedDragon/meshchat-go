package http

import (
	"embed"
	"fmt"

	"github.com/labstack/echo/v4"
)

//go:embed frontend/*
var fs embed.FS

// Server is the HTTP server.
type Server struct {
	*echo.Echo
	port int
}

// NewServer creates a new HTTP server.
func NewServer(port int) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.StaticFS("/", echo.MustSubFS(fs, "frontend"))
	return &Server{e, port}
}

// Start starts the HTTP server.
func (s *Server) Start() error {
	return s.Echo.Start(fmt.Sprintf(":%d", s.port))
}

// Stop stops the HTTP server.
func (s *Server) Stop() error {
	return s.Echo.Close()
}

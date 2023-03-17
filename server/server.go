package server

import (
	"fmt"

	"github.com/gaurishhs/uptimo/config"
	"github.com/gaurishhs/uptimo/database"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server struct {
	echo   *echo.Echo
	config *config.Config
	db     *database.Connection
}

const adminPath = "/_/"

func NewServer(dbConnection *database.Connection, config *config.Config) Server {
	server := Server{db: dbConnection, config: config}
	e := echo.New()
	e.Debug = config.Debug
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	if config.Debug {
		e.Use(middleware.Logger())
	}
	if config.Server.RateLimit > 0 {
		e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(config.Server.RateLimit)))
	}

	// Routes
	bindFrontend(e)

	server.echo = e
	return server
}

func (s *Server) Start() error {
	return s.echo.Start(fmt.Sprintf(":%d", s.config.Server.Port))
}

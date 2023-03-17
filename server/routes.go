package server

import (
	"net/http"
	"strings"

	"github.com/gaurishhs/uptimo/web"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// bindFrontend registers the frontend routes.
func bindFrontend(e *echo.Echo) error {
	// Make sure relative paths work
	e.GET(strings.TrimRight(adminPath, "/"), func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, strings.TrimLeft(adminPath, "/"))
	})

	// Serve Frontend with a Gzip middleware
	e.GET(adminPath+"*", echo.StaticDirectoryHandler(web.WebFS, false), middleware.Gzip())

	return nil
}

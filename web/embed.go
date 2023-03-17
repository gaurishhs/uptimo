package web

import (
	"embed"

	"github.com/labstack/echo/v5"
)

//go:embed all:dist
var embedWeb embed.FS

// WebFS is the filesystem containing the web assets.
var WebFS = echo.MustSubFS(embedWeb, "dist")

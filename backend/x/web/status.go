package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	StatusBadRequest          = echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	StatusInternalServerError = echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	StatusUnauthorized        = echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
)

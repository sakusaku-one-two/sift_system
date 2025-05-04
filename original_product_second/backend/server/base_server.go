package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var SERVER *echo.Echo = CreateServer()

func CreateServer() *echo.Echo {
	e := echo.New()
	return e
}

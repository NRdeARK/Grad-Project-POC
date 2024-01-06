package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Get_news(c echo.Context) error {
	return c.String(http.StatusOK,"this is news")
}
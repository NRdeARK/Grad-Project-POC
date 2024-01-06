package main

import (
    "github.com/labstack/echo/v4"

	"backend/internal/routes"
)

func main() {
    e := echo.New()
	e.GET("/news",routes.Get_news)
    e.Logger.Fatal(e.Start(":1323"))
}
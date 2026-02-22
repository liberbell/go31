package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func root(c echo.Context) error {
	return c.String(http.StatusOK, "Running API v1")
}

func main() {
	e := echo.New()

	e.GET("/", root)
	u := e.Group("/users")
	u.OPTIONS("", usersOptions)
	u.HEAD("", usersGetAll)

	e.Start(":12345")
}

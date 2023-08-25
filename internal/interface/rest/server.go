package rest

import (
	"net/http"

	"github.com/labstack/echo"
)

func NewServer() {
	e := echo.New()

	e.GET("/hello", helloWord)

	e.Logger.Fatal(e.Start(":3010"))
}

func helloWord(c echo.Context) error {
	resp := map[string]interface{}{
		"message": "hello word",
	}

	return c.JSON(http.StatusOK, resp)
}

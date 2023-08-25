package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Controllers struct {
	UploadController UploadController
}

func BindInputData(ctx context.Context, c echo.Context, data interface{}) error {
	if err := c.Bind(data); err != nil {
		return HandleError(http.StatusBadRequest, fmt.Sprintf("there is a bind error on your content: %s", err.Error()))
	}

	return nil
}

func (c *Controllers) NewServer() {
	e := echo.New()

	e.GET("/hello", helloWord)
	e.POST("/upload", c.UploadController.GetSign)

	e.Logger.Fatal(e.Start(":3010"))
}

func helloWord(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hello word",
	})
}

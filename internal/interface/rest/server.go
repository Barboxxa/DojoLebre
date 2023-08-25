package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/bavalogs.git"
	gotag_validator "gitlab.com/bavatech/architecture/software/libs/go-modules/gotag-validator.git/v2"
	"schneider.vip/problem"
)

type Controllers struct {
	UploadController UploadController
}

func BindInputData(ctx context.Context, c echo.Context, data interface{}) error {
	if err := c.Bind(data); err != nil {
		return HandleError(http.StatusBadRequest, fmt.Sprintf("there is a bind error on your content: %s", err.Error()))
	}

	if err := c.Validate(data); err.(*problem.Problem) != nil {
		return err
	}

	return nil
}

func (c *Controllers) NewServer() {

	validator, err := gotag_validator.NewValidator(nil, nil)
	if err != nil {
		bavalogs.Panic(context.TODO(), err).Msg("error trying to creating validator.")
	}

	e := echo.New()
	e.Validator = validator

	e.GET("/hello", helloWord)
	e.POST("/upload", c.UploadController.GetSign)

	e.Logger.Fatal(e.Start(":3010"))
}

func helloWord(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hello word",
	})
}

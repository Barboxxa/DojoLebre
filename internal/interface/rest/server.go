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

func BindInputData(ctx context.Context, c echo.Context, data interface{}) error {
	if err := c.Bind(data); err != nil {
		return HandleError(http.StatusBadRequest, fmt.Sprintf("there is a bind error on your content: %s", err.Error()))
	}

	if err := c.Validate(data); err.(*problem.Problem) != nil {
		return err
	}

	return nil
}

func NewServer() {

	validator, err := gotag_validator.NewValidator(nil, nil)
	if err != nil {
		bavalogs.Panic(context.TODO(), err).Msg("error trying to creating validator.")
	}

	e := echo.New()
	e.Validator = validator

	e.GET("/hello", helloWord)
	e.POST("/upload", upload)

	e.Logger.Fatal(e.Start(":3010"))
}

func helloWord(c echo.Context) error {
	resp := map[string]interface{}{
		"message": "hello word",
	}

	return c.JSON(http.StatusOK, resp)
}

type UploadRequest struct {
	Image string `json:"image" validate:"required,base64"`
}

func upload(c echo.Context) error {
	ctx := c.Request().Context()

	var payload UploadRequest

	if err := BindInputData(ctx, c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// TODO: deve fazer tratamentos e retornar a imagem base64;

	// TODO: Deve responder uma imagem base64;
	resp := map[string]interface{}{
		"image": "base64 image",
	}
	return c.JSON(http.StatusOK, resp)
}

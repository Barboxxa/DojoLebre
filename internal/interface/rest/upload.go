package rest

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/Barboxxa/DojoLebre/internal/domain"
	"github.com/Barboxxa/DojoLebre/internal/service"
)

type UploadController interface {
	GetSign(c echo.Context) error
}

type upload struct {
	uploadService service.Upload
}

func NewUploadController(uploadService service.Upload) UploadController {
	return &upload{
		uploadService,
	}
}

func (u upload) GetSign(c echo.Context) error {
	ctx := c.Request().Context()

	var payload domain.SignRequest

	if err := BindInputData(ctx, c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	encodedSign, err := u.uploadService.GetSign(ctx, payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// TODO: Deve responder uma imagem base64;
	resp := map[string]interface{}{
		"image": encodedSign,
	}
	return c.JSON(http.StatusOK, resp)
}

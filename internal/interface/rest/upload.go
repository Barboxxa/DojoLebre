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

	if err := validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	encodedSign, err := u.uploadService.GetSign(ctx, payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, HandleError(http.StatusInternalServerError, err.Error()))
	}

	resp := map[string]interface{}{
		"image": encodedSign,
	}
	return c.JSON(http.StatusOK, resp)
}

func validate(payload domain.SignRequest) error {
	// if _, err := base64.StdEncoding.DecodeString(payload.Image); err != nil {
	// 	return HandleError(http.StatusBadRequest, fmt.Sprintf("there is a validation error on body request: %s", err.Error()))
	// }

	return nil
}

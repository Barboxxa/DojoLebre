package rest

import (
	"net/http"

	"github.com/labstack/echo"
	"schneider.vip/problem"
)

func HandleError(httpStatus int, detail string, opts ...problem.Option) error {
	opts = append(opts, []problem.Option{
		problem.Status(httpStatus),
		problem.Title(http.StatusText(httpStatus)),
		problem.Detail(detail),
	}...)

	prob := problem.New(opts...)

	return echo.NewHTTPError(httpStatus, prob)
}

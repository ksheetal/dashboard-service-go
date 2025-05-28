package webserver

import (
	"context"
	"dashboard-service/internal/platform/config"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go.uber.org/fx"
	"net/http"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				NewWebServer,
				fx.OnStart(func(ctx context.Context, echo *echo.Echo, cfg *config.Config) error {
					err := setup(echo)
					if err != nil {
						return err
					}
					return nil
				}),
				fx.OnStop(func(echo *echo.Echo) error {
					return echo.Close()
				}),
			),
		))
}

type PlatformValidator struct {
	validator *validator.Validate
}

func (pv *PlatformValidator) Validate(i interface{}) error {
	if err := pv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
func newPlatformValidator() *PlatformValidator {
	v := validator.New()
	return &PlatformValidator{
		validator: v,
	}
}

func NewWebServer(cfg *config.Config) *echo.Echo {
	e := echo.New()
	return e
}

func setup(e *echo.Echo) error {
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Validator = newPlatformValidator()

	return nil

}

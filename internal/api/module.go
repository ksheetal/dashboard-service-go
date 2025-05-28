package api

import (
	"dashboard-service/internal/api/handler"
	"go.uber.org/fx"
)

func ApiModule() fx.Option {
	return fx.Options(
		handler.Module,
	)
}

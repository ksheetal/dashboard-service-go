package domain

import (
	"dashboard-service/internal/domain/dashboard"
	"go.uber.org/fx"
)

func DomainModule() fx.Option {
	return fx.Options(

		dashboard.DomainModule,
	)
}

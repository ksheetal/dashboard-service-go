package dashboard

import (
	"go.uber.org/fx"
)

var DomainModule = fx.Options(
	fx.Provide(
		NewReservation,
	),
)

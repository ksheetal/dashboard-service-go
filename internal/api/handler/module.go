package handler

import (
	"go.uber.org/fx"
)

// Module provides API handlers
var Module = fx.Options(
	fx.Provide(
		NewDashboardHandler,
	),
)

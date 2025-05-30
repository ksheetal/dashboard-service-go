package platform

import (
	"dashboard-service/internal/platform/file"
	"go.uber.org/fx"
)

func PlatformModule() fx.Option {
	return fx.Options(

		file.Module(),
	)
}

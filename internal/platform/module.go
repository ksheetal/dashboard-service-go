package platform

import (
	"dashboard-service/internal/platform/config"
	"dashboard-service/internal/platform/file"
	"dashboard-service/internal/platform/webserver"
	"go.uber.org/fx"
)

func PlatformModule() fx.Option {
	return fx.Options(
		config.Module(),
		webserver.Module(),
		file.Module(),
	)
}

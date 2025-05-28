package main

import (
	"context"
	"dashboard-service/internal/api"
	"dashboard-service/internal/domain"
	"dashboard-service/internal/platform"
	"dashboard-service/internal/platform/config"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Module("platform", platform.PlatformModule()),
		fx.Module("domain", domain.DomainModule()),
		fx.Module("api", api.ApiModule()),
		fx.Invoke(applicationLifeCycle),
	)
	app.Run()
}

func applicationLifeCycle(lc fx.Lifecycle, e *echo.Echo, cfg *config.Config) {

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				address := ":" + cfg.Platform.Server.Port
				err := e.Start(address)
				if err != nil {
					panic(err)
				}
				fmt.Println("Server is running on port " + address)

			}()
			return nil
		},
	})
}

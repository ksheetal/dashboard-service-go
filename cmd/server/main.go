package main

import (
	"context"
	"dashboard-service/internal/api"
	"dashboard-service/internal/domain"
	"dashboard-service/internal/domain/dashboard"
	"dashboard-service/internal/platform"
	"dashboard-service/internal/platform/config"
	"dashboard-service/internal/platform/file"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"time"
)

func main() {
	app := fx.New(
		fx.Module("platform", platform.PlatformModule()),
		fx.Module("domain", domain.DomainModule()),
		fx.Module("api", api.ApiModule()),
		fx.Invoke(applicationLifeCycle),
	)

	year := 2018
	monthNum := 01
	period := dashboard.Period{
		Year:  year,
		Month: time.Month(monthNum),
	}

	reservations, err := file.LoadReservations("cmd/server/metric_v1.csv")
	if err != nil {
		fmt.Println("Failed to load reservations:", err)
		return
	}

	service := dashboard.NewReservation()
	revenue, unreserved := service.Calculate(reservations, period)

	fmt.Printf("%04d-%02d: expected revenue: $%.2f, expected total capacity of the unreserved offices: %d\n",
		year, monthNum, revenue, unreserved)

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

/*
package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"

    "office-reservation/domain"
    "office-reservation/infrastructure"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run main.go <csv_file> <YYYY-MM>")
        return
    }

    filePath := os.Args[1]
    ym := strings.Split(os.Args[2], "-")
    year, _ := strconv.Atoi(ym[0])
    monthNum, _ := strconv.Atoi(ym[1])

    period := domain.Period{
        Year:  year,
        Month: time.Month(monthNum),
    }

    reservations, err := infrastructure.LoadReservations(filePath)
    if err != nil {
        fmt.Println("Failed to load reservations:", err)
        return
    }

    service := domain.NewReservationService()
    revenue, unreserved := service.Calculate(reservations, period)

    fmt.Printf("%04d-%02d: expected revenue: $%.2f, expected total capacity of the unreserved offices: %d\n",
        year, monthNum, revenue, unreserved)
}

*/

package main

import (
	"dashboard-service/internal/domain/dashboard"
	"dashboard-service/internal/platform/file"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

	period := dashboard.Period{
		Year:  year,
		Month: time.Month(monthNum),
	}

	fmt.Println(filePath)
	reservations, err := file.LoadReservations(filePath)
	if err != nil {
		fmt.Println("Failed to load reservations:", err)
		return
	}

	service := dashboard.NewReservation()
	revenue, unreserved := service.Calculate(reservations, period)

	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Printf("%04d-%02d: expected revenue: $%.2f, expected total capacity of the unreserved offices: %d\n",
		year, monthNum, revenue, unreserved)
	fmt.Println("---------------------------------------------------------------------------------------------")
}

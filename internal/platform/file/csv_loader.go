package file

import (
	"dashboard-service/internal/domain/dashboard"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func LoadReservations(filePath string) ([]dashboard.OfficeReservation, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var reservations []dashboard.OfficeReservation
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		capacity, _ := strconv.Atoi(strings.TrimSpace(record[0]))
		price, _ := strconv.ParseFloat(strings.TrimSpace(record[1]), 64)
		start, _ := time.Parse("2006-01-02", strings.TrimSpace(record[2]))
		var end *time.Time
		if strings.TrimSpace(record[3]) != "" {
			parsedEnd, _ := time.Parse("2006-01-02", strings.TrimSpace(record[3]))
			end = &parsedEnd
		}

		reservations = append(reservations, dashboard.OfficeReservation{
			Capacity:     capacity,
			MonthlyPrice: price,
			StartDate:    start,
			EndDate:      end,
		})
	}

	return reservations, nil
}

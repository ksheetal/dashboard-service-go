package dashboard

import (
	"time"
)

type OfficeReservation struct {
	Capacity     int
	MonthlyPrice float64
	StartDate    time.Time
	EndDate      *time.Time // nil means ongoing
}

type Period struct {
	Year  int
	Month time.Month
}

func (p Period) FirstDay() time.Time {
	return time.Date(p.Year, p.Month, 1, 0, 0, 0, 0, time.UTC)
}

func (p Period) LastDay() time.Time {
	return p.FirstDay().AddDate(0, 1, -1)
}

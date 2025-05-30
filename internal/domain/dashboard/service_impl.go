package dashboard

import "time"

type ReservationService struct {
}

func NewReservation() *ReservationService {
	return &ReservationService{}
}

func (s *ReservationService) Calculate(reservations []OfficeReservation, period Period) (float64, int) {

	revenue := 0.0
	totalReserved := make([]bool, len(reservations))

	for i, r := range reservations {
		start := r.StartDate
		end := r.EndDate
		monthStart := period.FirstDay()
		monthEnd := period.LastDay()

		// Determine overlap
		effectiveEnd := monthEnd
		if end != nil && end.Before(monthEnd) {
			effectiveEnd = *end
		}

		if start.After(monthEnd) || (end != nil && end.Before(monthStart)) {
			continue // No overlap
		}
		// Calculate prorated days
		overlapStart := maxDate(start, monthStart)
		overlapEnd := minDate(effectiveEnd, monthEnd)

		daysInMonth := monthEnd.Day()
		reservedDays := int(overlapEnd.Sub(overlapStart).Hours()/24) + 1

		if reservedDays > 0 {
			proratedRevenue := (r.MonthlyPrice / float64(daysInMonth)) * float64(reservedDays)
			revenue += proratedRevenue
			totalReserved[i] = true
		}
	}

	unreservedCapacity := 0
	for i, r := range reservations {
		if !totalReserved[i] {
			unreservedCapacity += r.Capacity
		}
	}

	return revenue, unreservedCapacity

}

func minDate(a, b time.Time) time.Time {
	if a.Before(b) {
		return a
	}
	return b
}

func maxDate(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}

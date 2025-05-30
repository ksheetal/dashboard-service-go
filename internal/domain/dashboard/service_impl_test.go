package dashboard

import (
	"testing"
	"time"
)

func TestCalculate(t *testing.T) {
	// Create a service instance
	service := NewReservation()

	// Define test cases
	tests := []struct {
		name         string
		reservations []OfficeReservation
		period       Period
		wantRevenue  float64
		wantCapacity int
	}{
		{
			name: "Full month reservation",
			reservations: []OfficeReservation{
				{
					Capacity:     10,
					MonthlyPrice: 5000,
					StartDate:    time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
					EndDate:      timePtr(time.Date(2018, 1, 31, 0, 0, 0, 0, time.UTC)),
				},
			},
			period:       Period{Year: 2018, Month: time.January},
			wantRevenue:  5000,
			wantCapacity: 0,
		},
		{
			name: "Partial month reservation",
			reservations: []OfficeReservation{
				{
					Capacity:     10,
					MonthlyPrice: 3100,
					StartDate:    time.Date(2018, 1, 16, 0, 0, 0, 0, time.UTC),
					EndDate:      timePtr(time.Date(2018, 1, 31, 0, 0, 0, 0, time.UTC)),
				},
			},
			period:       Period{Year: 2018, Month: time.January},
			wantRevenue:  1600, // ~(3100/31) * 16 days
			wantCapacity: 0,
		},
		{
			name: "Ongoing reservation (no end date)",
			reservations: []OfficeReservation{
				{
					Capacity:     5,
					MonthlyPrice: 2000,
					StartDate:    time.Date(2017, 12, 1, 0, 0, 0, 0, time.UTC),
					EndDate:      nil,
				},
			},
			period:       Period{Year: 2018, Month: time.January},
			wantRevenue:  2000,
			wantCapacity: 0,
		},
		{
			name: "No overlap - future reservation",
			reservations: []OfficeReservation{
				{
					Capacity:     8,
					MonthlyPrice: 4000,
					StartDate:    time.Date(2018, 2, 1, 0, 0, 0, 0, time.UTC),
					EndDate:      timePtr(time.Date(2018, 3, 31, 0, 0, 0, 0, time.UTC)),
				},
			},
			period:       Period{Year: 2018, Month: time.January},
			wantRevenue:  0,
			wantCapacity: 8,
		},
		{
			name: "No overlap - past reservation",
			reservations: []OfficeReservation{
				{
					Capacity:     4,
					MonthlyPrice: 1500,
					StartDate:    time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC),
					EndDate:      timePtr(time.Date(2017, 12, 31, 0, 0, 0, 0, time.UTC)),
				},
			},
			period:       Period{Year: 2018, Month: time.January},
			wantRevenue:  0,
			wantCapacity: 4,
		},
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			revenue, capacity := service.Calculate(tt.reservations, tt.period)

			// Compare revenue with a small epsilon to handle floating point precision
			if !floatEquals(revenue, tt.wantRevenue, 0.01) {
				t.Errorf("Calculate() revenue = %.2f, want %.2f", revenue, tt.wantRevenue)
			}

			if capacity != tt.wantCapacity {
				t.Errorf("Calculate() capacity = %d, want %d", capacity, tt.wantCapacity)
			}
		})
	}
}

// Helper function to create time pointer
func timePtr(t time.Time) *time.Time {
	return &t
}

// Helper function to compare floats with an epsilon
func floatEquals(a, b, epsilon float64) bool {
	return (a-b) < epsilon && (b-a) < epsilon
}

func TestPeriod_FirstDay(t *testing.T) {
	tests := []struct {
		name     string
		period   Period
		expected time.Time
	}{
		{
			name:     "January 2018",
			period:   Period{Year: 2018, Month: time.January},
			expected: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "February 2020 (leap year)",
			period:   Period{Year: 2020, Month: time.February},
			expected: time.Date(2020, 2, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "December 2025",
			period:   Period{Year: 2025, Month: time.December},
			expected: time.Date(2025, 12, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.period.FirstDay()
			if !result.Equal(tt.expected) {
				t.Errorf("FirstDay() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPeriod_LastDay(t *testing.T) {
	tests := []struct {
		name     string
		period   Period
		expected time.Time
	}{
		{
			name:     "January 2018 (31 days)",
			period:   Period{Year: 2018, Month: time.January},
			expected: time.Date(2018, 1, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "February 2020 (leap year - 29 days)",
			period:   Period{Year: 2020, Month: time.February},
			expected: time.Date(2020, 2, 29, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "February 2019 (non-leap year - 28 days)",
			period:   Period{Year: 2019, Month: time.February},
			expected: time.Date(2019, 2, 28, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "April 2018 (30 days)",
			period:   Period{Year: 2018, Month: time.April},
			expected: time.Date(2018, 4, 30, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.period.LastDay()
			if !result.Equal(tt.expected) {
				t.Errorf("LastDay() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestHelperFunctions(t *testing.T) {
	// Test minDate
	t.Run("minDate", func(t *testing.T) {
		t1 := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
		t2 := time.Date(2018, 1, 15, 0, 0, 0, 0, time.UTC)

		if res := minDate(t1, t2); !res.Equal(t1) {
			t.Errorf("minDate(%v, %v) = %v, want %v", t1, t2, res, t1)
		}

		if res := minDate(t2, t1); !res.Equal(t1) {
			t.Errorf("minDate(%v, %v) = %v, want %v", t2, t1, res, t1)
		}
	})

	// Test maxDate
	t.Run("maxDate", func(t *testing.T) {
		t1 := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
		t2 := time.Date(2018, 1, 15, 0, 0, 0, 0, time.UTC)

		if res := maxDate(t1, t2); !res.Equal(t2) {
			t.Errorf("maxDate(%v, %v) = %v, want %v", t1, t2, res, t2)
		}

		if res := maxDate(t2, t1); !res.Equal(t2) {
			t.Errorf("maxDate(%v, %v) = %v, want %v", t2, t1, res, t2)
		}
	})
}

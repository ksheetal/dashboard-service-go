package dashboard

import (
	"testing"
	"time"
)

func TestPeriodMethods(t *testing.T) {
	t.Run("Period with 31 days", func(t *testing.T) {
		period := Period{Year: 2018, Month: time.January}

		firstDay := period.FirstDay()
		expectedFirstDay := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
		if !firstDay.Equal(expectedFirstDay) {
			t.Errorf("FirstDay() = %v, want %v", firstDay, expectedFirstDay)
		}

		lastDay := period.LastDay()
		expectedLastDay := time.Date(2018, 1, 31, 0, 0, 0, 0, time.UTC)
		if !lastDay.Equal(expectedLastDay) {
			t.Errorf("LastDay() = %v, want %v", lastDay, expectedLastDay)
		}
	})

	t.Run("Period with 30 days", func(t *testing.T) {
		period := Period{Year: 2018, Month: time.April}

		firstDay := period.FirstDay()
		expectedFirstDay := time.Date(2018, 4, 1, 0, 0, 0, 0, time.UTC)
		if !firstDay.Equal(expectedFirstDay) {
			t.Errorf("FirstDay() = %v, want %v", firstDay, expectedFirstDay)
		}

		lastDay := period.LastDay()
		expectedLastDay := time.Date(2018, 4, 30, 0, 0, 0, 0, time.UTC)
		if !lastDay.Equal(expectedLastDay) {
			t.Errorf("LastDay() = %v, want %v", lastDay, expectedLastDay)
		}
	})

	t.Run("February in non-leap year", func(t *testing.T) {
		period := Period{Year: 2019, Month: time.February}

		firstDay := period.FirstDay()
		expectedFirstDay := time.Date(2019, 2, 1, 0, 0, 0, 0, time.UTC)
		if !firstDay.Equal(expectedFirstDay) {
			t.Errorf("FirstDay() = %v, want %v", firstDay, expectedFirstDay)
		}

		lastDay := period.LastDay()
		expectedLastDay := time.Date(2019, 2, 28, 0, 0, 0, 0, time.UTC)
		if !lastDay.Equal(expectedLastDay) {
			t.Errorf("LastDay() = %v, want %v", lastDay, expectedLastDay)
		}
	})

	t.Run("February in leap year", func(t *testing.T) {
		period := Period{Year: 2020, Month: time.February}

		firstDay := period.FirstDay()
		expectedFirstDay := time.Date(2020, 2, 1, 0, 0, 0, 0, time.UTC)
		if !firstDay.Equal(expectedFirstDay) {
			t.Errorf("FirstDay() = %v, want %v", firstDay, expectedFirstDay)
		}

		lastDay := period.LastDay()
		expectedLastDay := time.Date(2020, 2, 29, 0, 0, 0, 0, time.UTC)
		if !lastDay.Equal(expectedLastDay) {
			t.Errorf("LastDay() = %v, want %v", lastDay, expectedLastDay)
		}
	})
}

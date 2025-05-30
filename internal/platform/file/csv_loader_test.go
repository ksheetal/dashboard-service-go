package file

import (
	"dashboard-service/internal/domain/dashboard"
	"os"
	"testing"
	"time"
)

func TestLoadReservations(t *testing.T) {
	// Create a temporary test CSV file
	testCSV := `Capacity, Monthly Price, Start Day, End Day
10, 5000, 2018-01-01, 2018-03-31
5, 3000, 2018-02-15,
`

	tempFile, err := os.CreateTemp("", "test-reservations-*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up

	_, err = tempFile.Write([]byte(testCSV))
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	// Test loading the reservations
	reservations, err := LoadReservations(tempFile.Name())
	if err != nil {
		t.Fatalf("LoadReservations() error = %v", err)
	}

	// Check that we got the expected number of reservations
	if len(reservations) != 2 {
		t.Errorf("Expected 2 reservations, got %d", len(reservations))
	}

	// Check the first reservation's data
	expectedFirstReservation := dashboard.OfficeReservation{
		Capacity:     10,
		MonthlyPrice: 5000,
		StartDate:    time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:      timePtr(time.Date(2018, 3, 31, 0, 0, 0, 0, time.UTC)),
	}

	checkReservation(t, reservations[0], expectedFirstReservation)

	// Check the second reservation's data (with nil EndDate)
	expectedSecondReservation := dashboard.OfficeReservation{
		Capacity:     5,
		MonthlyPrice: 3000,
		StartDate:    time.Date(2018, 2, 15, 0, 0, 0, 0, time.UTC),
		EndDate:      nil, // ongoing reservation
	}

	checkReservation(t, reservations[1], expectedSecondReservation)
}

func TestLoadReservations_InvalidFile(t *testing.T) {
	// Test with non-existent file
	_, err := LoadReservations("non-existent-file.csv")
	if err == nil {
		t.Error("Expected error when loading non-existent file, got nil")
	}

	// Create a temporary test CSV file with invalid data
	testCSV := `Capacity, Monthly Price, Start Day, End Day
10, not-a-number, 2018-01-01, 2018-03-31
`

	tempFile, err := os.CreateTemp("", "test-invalid-*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up

	_, err = tempFile.Write([]byte(testCSV))
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	// Test loading the invalid reservations
	_, err = LoadReservations(tempFile.Name())
	if err == nil {
		t.Error("Expected error when loading invalid data, got nil")
	}
}

// Helper function to compare reservations
func checkReservation(t *testing.T, actual, expected dashboard.OfficeReservation) {
	t.Helper()

	if actual.Capacity != expected.Capacity {
		t.Errorf("Reservation Capacity = %d, want %d", actual.Capacity, expected.Capacity)
	}

	if actual.MonthlyPrice != expected.MonthlyPrice {
		t.Errorf("Reservation MonthlyPrice = %.2f, want %.2f", actual.MonthlyPrice, expected.MonthlyPrice)
	}

	if !actual.StartDate.Equal(expected.StartDate) {
		t.Errorf("Reservation StartDate = %v, want %v", actual.StartDate, expected.StartDate)
	}

	if (actual.EndDate == nil && expected.EndDate != nil) || 
	   (actual.EndDate != nil && expected.EndDate == nil) {
		t.Errorf("Reservation EndDate nil mismatch: got %v, want %v", actual.EndDate, expected.EndDate)
	} else if actual.EndDate != nil && expected.EndDate != nil && !actual.EndDate.Equal(*expected.EndDate) {
		t.Errorf("Reservation EndDate = %v, want %v", *actual.EndDate, *expected.EndDate)
	}
}

// Helper function to create time pointer
func timePtr(t time.Time) *time.Time {
	return &t
}

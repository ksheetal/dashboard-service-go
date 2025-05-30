# Office Reservation Dashboard Service

A Go-based service for calculating office reservation revenues and capacities for a specific time period.

## Project Overview

This dashboard service calculates expected revenue and available capacity for office reservations during a specified month. It processes reservation data from a CSV file and provides analytics based on the given time period.

### Key Features

- Calculate expected revenue for a specified month based on reservation data
- Determine total capacity of unreserved offices during that period
- Handle prorated calculations for partial month reservations

## Technology Stack

- **Programming Language**: Go 1.24.2
- **Architecture**: Domain-Driven Design pattern
- **Dependency Injection**: Uber FX

## Prerequisites

- Go 1.24.2 or later installed
- CSV file containing office reservation data

## Installation

1. Clone this repository
2. Navigate to the project directory
3. Install dependencies:
   ```
   go mod download
   ```

## How to Run

The application can be run using the following commands:

Move to server dir
```
cd cmd/server
```
then run 
```bash
go run main.go <csv_file> <YYYY-MM>
```

Where:
- `<csv_file>`: Path to the CSV file containing reservation data
- `<YYYY-MM>`: Year and month for which to calculate metrics (e.g., 2018-01)

### Example

```bash
go run main.go metric_v1 2018-01
```

### Expected Output

The application will display the expected revenue and unreserved office capacity for the specified month:

```
---------------------------------------------------------------------------------------------
2018-01: expected revenue: $1234.56, expected total capacity of the unreserved offices: 42
---------------------------------------------------------------------------------------------
```

## CSV File Format

The CSV file should contain the following columns:

1. Capacity - Number of people that can work in the office
2. Monthly Price - Cost of the office per month in dollars
3. Start Date - When the reservation begins (YYYY-MM-DD format)
4. End Date - When the reservation ends (YYYY-MM-DD format, can be empty for ongoing reservations)

### Example CSV Content

```csv
Capacity,MonthlyPrice,StartDate,EndDate
10,5000,2018-01-01,2018-03-15
5,3000,2018-02-01,
15,8000,2018-03-01,2018-05-31
```


## Testing

### Running Tests in a Specific Package

To run tests in a specific package (e.g., just the dashboard domain tests):

```bash
go test -v ./internal/domain/dashboard
```

### Running a Specific Test

To run a specific test function:

```bash
go test -v ./internal/domain/dashboard -run TestCalculate
```


#!/bin/bash

# Script to run tests for the Office Reservation Dashboard Service

set -e

# Print header
echo "========================================"
echo "  Running Office Reservation Tests"
echo "========================================"
echo ""

# Function to execute a test command and report status
run_test() {
  echo "📋 $1"
  echo "$ $2"
  eval $2
  echo "✅ Done"
  echo ""
}

# Run all tests
run_test "Running all tests" "go test ./..."

# Run tests with verbose output
run_test "Running tests with verbose output" "go test -v ./internal/domain/dashboard"

# Run test coverage
run_test "Generating test coverage report" "go test -coverprofile=coverage.out ./..."
echo "📊 Coverage summary:"
go tool cover -func=coverage.out
echo ""

# Generate HTML coverage report if browser is available
if command -v xdg-open >/dev/null 2>&1 || command -v open >/dev/null 2>&1; then
  echo "📈 Generating HTML coverage report"
  go tool cover -html=coverage.out -o coverage.html

  # Open the report in browser if possible
  if command -v xdg-open >/dev/null 2>&1; then
    xdg-open coverage.html
  elif command -v open >/dev/null 2>&1; then
    open coverage.html
  fi
fi

echo "========================================"
echo "  All tests completed"
echo "========================================"

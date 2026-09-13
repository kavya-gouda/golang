package main

import (
	"errors"
	"fmt"
)

// Tutorial demo 5: error handling
// Run: go run 05_errors.go

func main() {
	fmt.Println("=== (value, error) pattern ===")
	if result, err := divide(10, 0); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	fmt.Println("\n=== Sentinel error + errors.Is ===")
	_, err := withdraw(150, 100)
	if errors.Is(err, ErrInsufficientFunds) {
		fmt.Println("Handled a known error:", err)
	}

	fmt.Println("\n=== Custom error type + errors.As ===")
	err = validateAge(-5)
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Printf("Field %q had bad value %d\n", ve.Field, ve.Value)
	}

	fmt.Println("\n=== Wrapping with %w ===")
	err = loadUser()
	fmt.Println("Top-level error:", err)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Root cause was: not found")
	}
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %v by zero", a)
	}
	return a / b, nil
}

var ErrInsufficientFunds = errors.New("insufficient funds")

func withdraw(amount, balance int) (int, error) {
	if amount > balance {
		return balance, ErrInsufficientFunds
	}
	return balance - amount, nil
}

type ValidationError struct {
	Field string
	Value int
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid %s: %d", e.Field, e.Value)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Value: age}
	}
	return nil
}

var ErrNotFound = errors.New("not found")

func loadUser() error {
	return fmt.Errorf("loadUser: could not fetch user: %w", ErrNotFound)
}

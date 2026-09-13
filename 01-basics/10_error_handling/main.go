package main

import (
	"errors"
	"fmt"
)

// Lesson: Error handling in Go
//
// Go does not use try/catch for normal errors. Instead, functions
// return an error value that you check explicitly. This makes error
// paths visible and hard to ignore.
//
// Run this file:
//   go run main.go

func main() {
	fmt.Println("=== 1. The standard error pattern ===")
	// A function that can fail returns (result, error).
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // this branch runs
	} else {
		fmt.Println("10 / 0 =", result)
	}

	fmt.Println("\n=== 2. Sentinel errors + errors.Is ===")
	// Try to take out 150 when only 100 is available -> triggers the error.
	_, err = withdraw(150, 100)
	if errors.Is(err, ErrInsufficientFunds) {
		fmt.Println("Handled a known error:", err)
	}

	fmt.Println("\n=== 3. Custom error type ===")
	err = validateAge(-5)
	if err != nil {
		fmt.Println("Validation failed:", err)
	}

	// errors.As lets you inspect the concrete error type.
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Printf("Field %q had bad value %d\n", ve.Field, ve.Value)
	}

	fmt.Println("\n=== 4. Wrapping errors with %w ===")
	err = loadUser()
	fmt.Println("Top-level error:", err)
	// Because we wrapped it, we can still detect the original cause:
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Root cause was: not found")
	}

	// -------------------------------------------------------------------
	// YOUR TURN:
	// 1. Add a multiply function that never fails. Does it need an error?
	// 2. Call withdraw(150, 100). Which error comes back?
	// 3. Make validateAge also reject ages over 150 with a clear message.
	// -------------------------------------------------------------------
}

// divide returns an error instead of crashing on divide-by-zero.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %v by zero", a)
	}
	return a / b, nil
}

// Sentinel error: a predefined error value you can compare against.
var ErrInsufficientFunds = errors.New("insufficient funds")

func withdraw(amount, balance int) (int, error) {
	if amount > balance {
		return balance, ErrInsufficientFunds
	}
	return balance - amount, nil
}

// Custom error type: carries extra structured data.
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

// Wrapping: add context while preserving the original error.
var ErrNotFound = errors.New("not found")

func loadUser() error {
	// Imagine a database lookup failed with ErrNotFound.
	return fmt.Errorf("loadUser: could not fetch user: %w", ErrNotFound)
}

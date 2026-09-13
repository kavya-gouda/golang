package main

import "fmt"

// Tutorial demo 3: switch
// Run: go run 03_switch.go

func main() {
	fmt.Println("=== Value switch (no fall-through, multiple values) ===")
	day := "Saturday"
	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("A regular day")
	}

	fmt.Println("\n=== Conditionless switch (if-else replacement) ===")
	score := 72
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

	fmt.Println("\n=== fallthrough (opt-in) ===")
	num := 1
	switch num {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println("\n=== type switch ===")
	describe(42)
	describe("hello")
	describe(true)
}

func describe(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Printf("int: %d\n", v)
	case string:
		fmt.Printf("string of length %d: %q\n", len(v), v)
	case bool:
		fmt.Printf("bool: %v\n", v)
	default:
		fmt.Printf("unknown type %T\n", v)
	}
}

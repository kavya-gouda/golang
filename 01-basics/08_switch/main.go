package main

import "fmt"

// Lesson: switch statements in Go
//
// switch is a clean alternative to long if-else-if chains.
// Key difference from C/Java: Go does NOT fall through by default,
// so you don't need `break` at the end of each case.
//
// Run this file:
//   go run main.go

func main() {
	fmt.Println("=== 1. Basic switch on a value ===")
	day := "Saturday"
	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Saturday", "Sunday": // multiple values in one case
		fmt.Println("Weekend!")
	default:
		fmt.Println("A regular day")
	}

	fmt.Println("\n=== 2. switch with no condition (replaces if-else chain) ===")
	score := 72
	switch { // no value after switch = evaluate boolean cases
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

	fmt.Println("\n=== 3. switch with a short statement ===")
	switch hour := 14; {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	fmt.Println("\n=== 4. fallthrough (opt-in to next case) ===")
	num := 1
	switch num {
	case 1:
		fmt.Println("one")
		fallthrough // force execution of the NEXT case
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	// Output: one, two  (fallthrough ran case 2, but stopped there)

	fmt.Println("\n=== 5. type switch ===")
	describe(42)
	describe("hello")
	describe(true)
	describe(3.14)

	// -------------------------------------------------------------------
	// YOUR TURN:
	// 1. Write a switch that maps a month number (1-12) to a season.
	// 2. Change `num` in example 4 to 2. What prints and why?
	// 3. Add a `case []int:` to describe() and pass a slice to it.
	// -------------------------------------------------------------------
}

// type switch: react based on the underlying type of an interface value
func describe(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Printf("int: %d\n", v)
	case string:
		fmt.Printf("string of length %d: %q\n", len(v), v)
	case bool:
		fmt.Printf("bool: %v\n", v)
	default:
		fmt.Printf("unknown type %T with value %v\n", v, v)
	}
}

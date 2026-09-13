package main

import "fmt"

// Lesson: if / else / else-if in Go
//
// Run this file:
//   go run main.go
//
// Then try the "Your Turn" challenges at the bottom by editing values.

func main() {
	fmt.Println("=== 1. Basic if ===")
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	fmt.Println("\n=== 2. if / else ===")
	age := 16
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	fmt.Println("\n=== 3. if / else-if / else ===")
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	fmt.Println("\n=== 4. if with a short statement ===")
	// You can declare a variable that only lives inside the if/else.
	// This is one of Go's most-used patterns, especially for errors.
	if n := len("hello"); n > 3 {
		fmt.Println("The word is long, length =", n)
	} else {
		fmt.Println("The word is short, length =", n)
	}
	// Note: n does NOT exist out here. Uncommenting the next line is a compile error:
	// fmt.Println(n)

	fmt.Println("\n=== 5. Logical operators (&& || !) ===")
	hasLicense := true
	if age >= 18 && hasLicense {
		fmt.Println("Can drive")
	}
	if age < 18 || !hasLicense {
		fmt.Println("Cannot drive")
	}

	// -------------------------------------------------------------------
	// YOUR TURN (edit and re-run):
	// 1. Change `age` to 20 and `hasLicense` to false. Who can drive now?
	// 2. Change `score` and predict the grade before running.
	// 3. Write an if-statement that prints "even" or "odd" for a number.
	//    Hint: use the modulo operator, e.g. num%2 == 0
	// -------------------------------------------------------------------
}

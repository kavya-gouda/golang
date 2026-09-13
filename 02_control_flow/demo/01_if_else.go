package main

import "fmt"

// Tutorial demo 1: if / else
// Run: go run 01_if_else.go

func main() {
	fmt.Println("=== Basic if ===")
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	fmt.Println("\n=== if / else-if / else ===")
	score := 85 // try changing to 95 on camera
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	fmt.Println("\n=== Short-statement if (the error pattern) ===")
	if n := len("hello"); n > 3 {
		fmt.Println("word is long, length =", n)
	}
	// The idiomatic error check looks like this:
	//   if err := doThing(); err != nil {
	//       return err
	//   }

	fmt.Println("\n=== Logical operators && || ! ===")
	age := 20
	hasLicense := true
	if age >= 18 && hasLicense {
		fmt.Println("Can drive")
	}
	if age < 18 || !hasLicense {
		fmt.Println("Cannot drive")
	}
}

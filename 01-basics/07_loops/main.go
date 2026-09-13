package main

import "fmt"

// Lesson: Loops in Go
//
// Go has ONLY ONE loop keyword: `for`.
// It can act like a classic for-loop, a while-loop, and an infinite loop.
//
// Run this file:
//   go run main.go

func main() {
	fmt.Println("=== 1. Classic for loop ===")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ") // 0 1 2 3 4
	}
	fmt.Println()

	fmt.Println("\n=== 2. 'while' style loop ===")
	// No 'while' keyword. Just drop the init and post parts.
	count := 0
	for count < 3 {
		fmt.Println("count is", count)
		count++
	}

	fmt.Println("\n=== 3. Infinite loop with break ===")
	i := 0
	for { // no condition = runs forever until break
		if i == 3 {
			break // exit the loop
		}
		fmt.Println("tick", i)
		i++
	}

	fmt.Println("\n=== 4. continue (skip an iteration) ===")
	for n := 0; n < 6; n++ {
		if n%2 != 0 {
			continue // skip odd numbers
		}
		fmt.Println("even:", n)
	}

	fmt.Println("\n=== 5. range over a slice ===")
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("%d -> %s\n", index, fruit)
	}

	fmt.Println("\n=== 6. range over a map ===")
	ages := map[string]int{"Alice": 25, "Bob": 30}
	for name, age := range ages {
		fmt.Printf("%s is %d\n", name, age)
	}

	fmt.Println("\n=== 7. range over a string (runes) ===")
	for i, ch := range "Go!" {
		fmt.Printf("index %d = %c\n", i, ch)
	}

	fmt.Println("\n=== 8. Ignoring index or value with _ ===")
	sum := 0
	for _, v := range []int{10, 20, 30} {
		sum += v // we only care about the value
	}
	fmt.Println("sum =", sum)

	fmt.Println("\n=== 9. Labeled break (exit nested loops) ===")
outer:
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if r == 1 && c == 1 {
				break outer // breaks BOTH loops
			}
			fmt.Printf("r=%d c=%d\n", r, c)
		}
	}

	// -------------------------------------------------------------------
	// YOUR TURN:
	// 1. Print numbers 10 down to 1 (a countdown). Hint: i-- and i > 0.
	// 2. Sum all numbers in []int{5, 7, 9, 11} using range.
	// 3. Loop over a string and count how many vowels it has.
	// -------------------------------------------------------------------
}

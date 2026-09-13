package main

import "fmt"

// Tutorial demo 2: loops (the one 'for' keyword)
// Run: go run 02_loops.go

func main() {
	fmt.Println("=== Classic for ===")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	fmt.Println("\n=== While-style ===")
	count := 0
	for count < 3 {
		fmt.Println("count", count)
		count++
	}

	fmt.Println("\n=== Infinite loop + break ===")
	i := 0
	for {
		if i == 3 {
			break
		}
		fmt.Println("tick", i)
		i++
	}

	fmt.Println("\n=== range: slice ===")
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("%d -> %s\n", index, fruit)
	}

	fmt.Println("\n=== range: map (order is randomized!) ===")
	ages := map[string]int{"Alice": 25, "Bob": 30, "Carol": 28}
	for name, age := range ages {
		fmt.Printf("%s is %d\n", name, age)
	}

	fmt.Println("\n=== continue (skip odds) ===")
	for n := 0; n < 6; n++ {
		if n%2 != 0 {
			continue
		}
		fmt.Println("even:", n)
	}

	fmt.Println("\n=== labeled break (exit both loops) ===")
outer:
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if r == 1 && c == 1 {
				break outer
			}
			fmt.Printf("r=%d c=%d\n", r, c)
		}
	}
}

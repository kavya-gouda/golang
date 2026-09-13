package main

import "fmt"

// Lesson: Practice combining if, for, and switch
//
// These are worked solutions to the exercises in the control flow guide.
// Try to solve each one yourself first, then compare with these.
//
// Run this file:
//   go run main.go

func main() {
	fmt.Println("=== Exercise 1: FizzBuzz (1..20) ===")
	fizzBuzz(20)

	fmt.Println("\n=== Exercise 2: Prime checker ===")
	for _, n := range []int{1, 2, 7, 9, 13, 15, 29} {
		fmt.Printf("isPrime(%d) = %v\n", n, isPrime(n))
	}

	fmt.Println("\n=== Exercise 3: Calculator ===")
	printCalc(6, 3, "+")
	printCalc(6, 3, "-")
	printCalc(6, 3, "*")
	printCalc(6, 3, "/")
	printCalc(6, 0, "/") // division by zero handled

	fmt.Println("\n=== Exercise 4: Triangle pattern ===")
	triangle(5)
}

// fizzBuzz prints 1..n, replacing multiples of 3/5 with Fizz/Buzz.
func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0: // divisible by both 3 and 5
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

// isPrime returns true if n is a prime number.
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ { // only check up to sqrt(n)
		if n%i == 0 {
			return false
		}
	}
	return true
}

// calculate performs a basic operation and reports an error for /0.
func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operator %q", op)
	}
}

func printCalc(a, b float64, op string) {
	result, err := calculate(a, b, op)
	if err != nil {
		fmt.Printf("%g %s %g -> error: %v\n", a, op, b, err)
		return
	}
	fmt.Printf("%g %s %g = %g\n", a, op, b, result)
}

// triangle prints a right triangle of stars with the given height.
func triangle(height int) {
	for row := 1; row <= height; row++ {
		for star := 0; star < row; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

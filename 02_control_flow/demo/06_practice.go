package main

import "fmt"

// Tutorial demo 6: everything together
// Run: go run 06_practice.go

func main() {
	fmt.Println("=== FizzBuzz (1..20) ===")
	fizzBuzz(20)

	fmt.Println("\n=== Prime checker ===")
	for _, n := range []int{2, 9, 13, 15, 29} {
		fmt.Printf("isPrime(%d) = %v\n", n, isPrime(n))
	}

	fmt.Println("\n=== Calculator ===")
	printCalc(6, 3, "+")
	printCalc(6, 3, "/")
	printCalc(6, 0, "/")

	fmt.Println("\n=== Triangle ===")
	triangle(5)
}

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
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

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

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

func triangle(height int) {
	for row := 1; row <= height; row++ {
		for star := 0; star < row; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

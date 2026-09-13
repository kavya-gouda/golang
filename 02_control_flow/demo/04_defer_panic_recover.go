package main

import "fmt"

// Tutorial demo 4: defer, panic, recover
// Run: go run 04_defer_panic_recover.go

func main() {
	fmt.Println("=== defer runs LIFO ===")
	deferDemo()

	fmt.Println("\n=== defer for cleanup ===")
	cleanupDemo()

	fmt.Println("\n=== panic + recover ===")
	fmt.Println("Before safeDivide")
	fmt.Println("10 / 0 =", safeDivide(10, 0))
	fmt.Println("Program continues because we recovered!")
}

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("deferred A (runs last)")
	defer fmt.Println("deferred B (runs first)")
	fmt.Println("end")
}

func cleanupDemo() {
	fmt.Println("opening resource...")
	defer fmt.Println("closing resource (always runs)")
	fmt.Println("using resource...")
}

func safeDivide(a, b int) (answer int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from:", r)
			answer = -1
		}
	}()
	return a / b
}

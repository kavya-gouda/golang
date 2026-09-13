package main

import "fmt"

// Lesson: defer, panic, and recover
//
// defer   -> schedule a call to run when the surrounding function returns
// panic   -> stop normal flow (like throwing an exception)
// recover -> catch a panic inside a deferred function
//
// Run this file:
//   go run main.go

func main() {
	fmt.Println("=== 1. defer runs at the end (LIFO order) ===")
	deferDemo()

	fmt.Println("\n=== 2. defer is great for cleanup ===")
	cleanupDemo()

	fmt.Println("\n=== 3. panic + recover ===")
	fmt.Println("Before calling safeDivide")
	result := safeDivide(10, 0)
	fmt.Println("Result:", result)
	fmt.Println("Program keeps running because we recovered!")

	// -------------------------------------------------------------------
	// YOUR TURN:
	// 1. Add a third defer to deferDemo(). Predict the print order first.
	// 2. Call safeDivide(10, 2). Does it still work without panicking?
	// 3. Remove the recover() in safeDivide and run again to see a crash.
	// -------------------------------------------------------------------
}

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("deferred A (runs last)")
	defer fmt.Println("deferred B (runs before A)")
	fmt.Println("end")
	// Deferred calls run in Last-In-First-Out order:
	// start, end, deferred B, deferred A
}

func cleanupDemo() {
	fmt.Println("opening resource...")
	defer fmt.Println("closing resource (always runs)")

	fmt.Println("using resource...")
	// Even if we returned early or panicked, the deferred close still runs.
}

// safeDivide uses recover to turn a panic into a safe return value.
func safeDivide(a, b int) (answer int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from:", r)
			answer = -1 // provide a fallback result
		}
	}()

	return a / b // dividing by zero triggers a panic
}

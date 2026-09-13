# Control Flow and Logic in Go - Complete Guide

## Table of Contents
1. [Conditional Statements](#conditional-statements)
2. [Loops](#loops)
3. [Switch Statements](#switch-statements)
4. [Defer, Panic, and Recover](#defer-panic-recover)
5. [Error Handling](#error-handling)
6. [Control Flow Patterns](#control-flow-patterns)
7. [Practice Exercises](#practice-exercises)

---

## Conditional Statements

### 1. If Statement

#### Basic If
```go
package main

import "fmt"

func main() {
    x := 10
    
    if x > 5 {
        fmt.Println("x is greater than 5")
    }
}
```

#### If-Else
```go
age := 18

if age >= 18 {
    fmt.Println("You are an adult")
} else {
    fmt.Println("You are a minor")
}
```

#### If-Else If-Else
```go
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
```

#### If with Short Statement
```go
// Initialize and check in one line
if x := 10; x > 5 {
    fmt.Println("x is greater than 5")
}
// x is only available inside the if block

// Common pattern with error checking
if err := someFunction(); err != nil {
    fmt.Println("Error:", err)
}
```

### 2. Logical Operators

```go
a := true
b := false

// AND (&&)
if a && b {
    fmt.Println("Both are true")
}

// OR (||)
if a || b {
    fmt.Println("At least one is true")
}

// NOT (!)
if !b {
    fmt.Println("b is false")
}

// Combining operators
age := 25
hasLicense := true

if age >= 18 && hasLicense {
    fmt.Println("Can drive")
}

if age < 18 || !hasLicense {
    fmt.Println("Cannot drive")
}
```

---

## Loops

### 1. For Loop (Go's only loop)

#### Basic For Loop
```go
// Traditional for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)  // 0, 1, 2, 3, 4
}
```

#### While-style Loop
```go
count := 0

// While loop (Go style)
for count < 5 {
    fmt.Println(count)
    count++
}
```

#### Infinite Loop
```go
// Infinite loop (use with caution!)
for {
    fmt.Println("This runs forever")
    // Use break to exit
    break
}

// Practical infinite loop
for {
    // Read user input
    // Process data
    // Break on certain condition
    if someCondition {
        break
    }
}
```

#### For-Range Loop (for collections)
```go
// Array
numbers := [3]int{10, 20, 30}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// Slice
fruits := []string{"apple", "banana", "cherry"}
for i, fruit := range fruits {
    fmt.Printf("%d: %s\n", i, fruit)
}

// Map
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 28,
}
for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}

// String (iterates over runes, not bytes)
for i, ch := range "Hello" {
    fmt.Printf("Index: %d, Character: %c\n", i, ch)
}

// Ignoring index/value
for _, value := range numbers {
    fmt.Println(value)  // Only values, no index
}

for index := range numbers {
    fmt.Println(index)  // Only index, no values
}
```

### 2. Loop Control Statements

#### Break
```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break  // Exit loop completely
    }
    fmt.Println(i)  // 0, 1, 2, 3, 4
}
```

#### Continue
```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue  // Skip this iteration
    }
    fmt.Println(i)  // 0, 1, 3, 4
}
```

#### Labeled Break/Continue
```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer  // Break out of both loops
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```

---

## Switch Statements

### 1. Basic Switch
```go
day := "Monday"

switch day {
case "Monday":
    fmt.Println("Start of work week")
case "Tuesday", "Wednesday", "Thursday":
    fmt.Println("Mid week")
case "Friday":
    fmt.Println("Almost weekend!")
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
default:
    fmt.Println("Invalid day")
}
```

### 2. Switch with Expressions
```go
score := 85

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
```

### 3. Type Switch
```go
func printType(x interface{}) {
    switch v := x.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %v\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}

printType(42)      // Integer: 42
printType("hello") // String: hello
printType(true)    // Boolean: true
```

### 4. Fallthrough
```go
// Go switch doesn't fall through by default
num := 2

switch num {
case 1:
    fmt.Println("One")
    fallthrough  // Continue to next case
case 2:
    fmt.Println("Two")
    fallthrough
case 3:
    fmt.Println("Three")
}
// Output: Two, Three
```

---

## Defer, Panic, and Recover

### 1. Defer
```go
func main() {
    fmt.Println("Start")
    
    defer fmt.Println("This runs last (deferred)")
    defer fmt.Println("This runs second (LIFO - Last In First Out)")
    
    fmt.Println("End")
}
// Output: Start, End, This runs second, This runs last
```

### 2. Panic
```go
func riskyFunction() {
    panic("Something went wrong!")
}

func main() {
    fmt.Println("Before panic")
    riskyFunction()
    fmt.Println("After panic")  // Never reached
}
```

### 3. Recover
```go
func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    
    panic("Intentional panic")
    fmt.Println("This won't execute")
}

func main() {
    fmt.Println("Before safe function")
    safeFunction()
    fmt.Println("After safe function")  // This executes!
}
```

---

## Error Handling

### 1. Basic Error Handling
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
    
    // Try with zero
    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)  // This executes
    }
}
```

### 2. Custom Error Types
```go
type DivisionError struct {
    Numerator   float64
    Denominator float64
    Message     string
}

func (e *DivisionError) Error() string {
    return fmt.Sprintf("%s: %.2f / %.2f", e.Message, e.Numerator, e.Denominator)
}

func safeDivide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, &DivisionError{
            Numerator:   a,
            Denominator: b,
            Message:     "division by zero",
        }
    }
    return a / b, nil
}
```

---

## Control Flow Patterns

### 1. Early Return Pattern
```go
func processUser(user string, age int) error {
    // Validate early
    if user == "" {
        return fmt.Errorf("username cannot be empty")
    }
    
    if age < 0 {
        return fmt.Errorf("age cannot be negative")
    }
    
    if age < 18 {
        return fmt.Errorf("user must be at least 18 years old")
    }
    
    // Main logic
    fmt.Printf("Processing user: %s, age: %d\n", user, age)
    return nil
}
```

### 2. Loop with Break Pattern
```go
func findNumber(numbers []int, target int) (bool, int) {
    for i, num := range numbers {
        if num == target {
            return true, i  // Early return when found
        }
    }
    return false, -1  // Not found
}
```

### 3. Nested Loop Pattern
```go
func findPairs(numbers []int, targetSum int) [][2]int {
    var pairs [][2]int
    
    for i := 0; i < len(numbers); i++ {
        for j := i + 1; j < len(numbers); j++ {
            if numbers[i]+numbers[j] == targetSum {
                pairs = append(pairs, [2]int{numbers[i], numbers[j]})
            }
        }
    }
    return pairs
}
```

### 4. Defer for Cleanup Pattern
```go
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // Always close the file
    
    // Process file
    // ...
    return nil
}
```

### 5. Switch for State Machine
```go
type State int

const (
    Idle State = iota
    Running
    Paused
    Stopped
)

func handleState(state State) {
    switch state {
    case Idle:
        fmt.Println("System is idle")
    case Running:
        fmt.Println("System is running")
    case Paused:
        fmt.Println("System is paused")
    case Stopped:
        fmt.Println("System is stopped")
    }
}
```

---

## Practice Exercises

### Exercise 1: FizzBuzz
```go
// Print numbers 1 to 100
// For multiples of 3, print "Fizz"
// For multiples of 5, print "Buzz"
// For multiples of both, print "FizzBuzz"
```

### Exercise 2: Prime Number Checker
```go
// Write a function that checks if a number is prime
func isPrime(n int) bool {
    // Your code here
}
```

### Exercise 3: Calculator
```go
// Create a calculator that takes two numbers and an operator
// Use switch statement to handle +, -, *, / operations
// Handle division by zero error
```

### Exercise 4: Pattern Printer
```go
// Print patterns using nested loops
// Example: Right triangle
// *
// **
// ***
// ****
// *****
```

### Exercise 5: Guessing Game
```go
// Create a number guessing game
// Generate random number between 1-100
// Give user 5 attempts
// Provide hints (too high/too low)
```

---

## Common Pitfalls

### 1. Infinite Loops
```go
// Wrong
for i := 0; i < 5; i-- {  // i-- instead of i++
    fmt.Println(i)  // Infinite loop!
}

// Right
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### 2. Off-by-One Errors
```go
// Wrong
for i := 0; i <= len(slice); i++ {  // <= instead of <
    // Index out of bounds!
}

// Right
for i := 0; i < len(slice); i++ {
    // Safe
}
```

### 3. Missing Break in Switch
```go
// Go doesn't need break (unlike other languages)
// This is correct in Go:
switch x {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
}
```

### 4. Defer Timing
```go
func wrongDefer() {
    for i := 0; i < 3; i++ {
        defer fmt.Println(i)  // All print 2! (deferred value)
    }
}

func rightDefer() {
    for i := 0; i < 3; i++ {
        j := i  // Capture loop variable
        defer fmt.Println(j)  // Prints 2, 1, 0
    }
}
```

---

## Best Practices

1. **Use early returns** to reduce nesting
2. **Keep loops simple** - extract complex logic into functions
3. **Use `defer` for cleanup** (files, connections, etc.)
4. **Prefer `switch` over long `if-else` chains**
5. **Handle errors immediately** after function calls
6. **Use labeled breaks** for nested loops instead of flags
7. **Keep conditionals simple** - extract complex conditions into variables
8. **Use `range`** for iterating over collections
9. **Avoid `goto`** (Go has it but rarely needed)
10. **Write testable control flow** - make functions small and focused

---

## Summary

- **If statements**: Basic conditionals with optional short statements
- **For loops**: Go's only loop (can mimic while, do-while, infinite)
- **Switch statements**: Clean alternative to long if-else chains
- **Defer**: Schedule cleanup operations
- **Panic/Recover**: Handle exceptional situations
- **Error handling**: Use multiple return values for errors

Practice these patterns to become proficient in Go's control flow! Start with simple exercises and gradually tackle more complex problems.
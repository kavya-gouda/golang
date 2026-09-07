# Enums in Go - In-Depth Guide

## Overview

Go doesn't have a dedicated `enum` keyword like C++, Java, or TypeScript. Instead, Go uses **constants with `iota`** to create enumerated types. This approach is idiomatic and provides type safety when combined with custom types.

---

## What is an Enum? (Beginner's Guide)

An **enum** (short for "enumeration") is a way to create a list of named constants in programming. Think of it like creating a menu with fixed options.

### Everyday Example: Days of the Week

Imagine you're writing a program that needs to track days of the week. You could do this:

```go
// Without enum - using magic numbers
scheduleTask(1)  // Monday?
scheduleTask(3)  // Wednesday?
```

This is confusing! What does `1` mean? What does `3` mean? You have to remember which number maps to which day.

### The Better Way: Using Names

Instead of numbers, use meaningful names:

```go
// With enum-like approach
scheduleTask("Monday")
scheduleTask("Wednesday")
```

Much clearer! But what if someone makes a typo?

```go
scheduleTask("Mondya")  // Oops, typo!
```

The program will fail at runtime, not at compile time.

### The Best Way: Enums

Enums solve both problems:
1. **Use meaningful names** (not magic numbers)
2. **Catch errors at compile time** (not runtime)

---

## Why Use Enums? Simple Benefits

### 1. **Readability**
```go
// Bad: Magic numbers
if status == 2 { /* ??? */ }

// Good: Clear meaning
if status == Active { /* Much clearer! */ }
```

### 2. **Type Safety**
```go
type Status int
const (
    Pending Status = iota
    Active
    Completed
)

func updateStatus(s Status) {
    // Can only accept Status values
}

updateStatus(Active)    // ✅ Works
updateStatus(999)       // ❌ Error! Can't pass random number
updateStatus("Active")  // ❌ Error! Can't pass string
```

### 3. **Prevents Typos**
```go
type Color string
const (
    Red   Color = "red"
    Green Color = "green"
    Blue  Color = "blue"
)

func setColor(c Color) {
    // Only accepts our defined colors
}

setColor(Red)           // ✅ Works
setColor("red")         // ✅ Works (but less safe)
setColor("purple")      // ❌ Compile error (if using type)
setColor("redd")        // ❌ Runtime error (typo!)
```

---

## Real-World Simple Examples for Beginners

### Example 1: Pizza Size
```go
type PizzaSize int

const (
    Small PizzaSize = iota
    Medium
    Large
    XLarge
)

func orderPizza(size PizzaSize) {
    fmt.Printf("Ordering %s pizza\n", size)
}
```

### Example 2: Traffic Light
```go
type TrafficLight int

const (
    Red TrafficLight = iota
    Yellow
    Green
)

func drive(light TrafficLight) {
    switch light {
    case Red:
        fmt.Println("Stop!")
    case Yellow:
        fmt.Println("Slow down...")
    case Green:
        fmt.Println("Go!")
    }
}
```

### Example 3: Game Difficulty
```go
type Difficulty int

const (
    Easy Difficulty = iota
    Medium
    Hard
)

func startGame(difficulty Difficulty) {
    fmt.Printf("Starting %s game\n", difficulty)
}
```

---

## How Enums Help You

### Before Enums (Confusing)
```go
// What does 1 mean? What does 2 mean?
var status = 1
if status == 2 {
    // ??? 
}
```

### After Enums (Clear)
```go
type TaskStatus int
const (
    Pending TaskStatus = iota
    InProgress
    Completed
)

var status = InProgress
if status == Completed {
    // Very clear!
}
```

---

## Quick Practice Exercise

Here's a mini-exercise for you. Try to understand what this code does:

```go
type Size int

const (
    Tiny Size = iota
    Small
    Medium
    Large
    Huge
)

func printClothingSize(size Size) {
    sizes := []string{"Tiny", "Small", "Medium", "Large", "Huge"}
    fmt.Printf("You need size: %s\n", sizes[size])
}

func main() {
    mySize := Medium
    printClothingSize(mySize)  // What will this print?
}
```

**Answer:** It prints "You need size: Medium" because `Medium` has value `2`, and `sizes[2]` is "Medium".

---

## Common Beginner Questions

**Q: Why not just use strings?**
```go
// Bad: Strings can have typos
var color = "redd"  // Typo!

// Good: Enum catches this
var color = Color.Red  // Always correct
```

**Q: Why not just use numbers?**
```go
// Bad: Magic numbers
if day == 1 { /* Monday? Sunday? */ }

// Good: Clear names
if day == Monday { /* Obvious! */ }
```

**Q: When should I use enums?**
- When you have a fixed set of options
- When those options don't change often
- When you want to prevent wrong values
- When you want readable code

---

## Bottom Line for Beginners

Enums are like creating a **menu** for your code:
- You define what's on the menu (allowed values)
- Users can only order from the menu (no typos)
- Everyone knows what they're getting (readable)
- The kitchen knows what to prepare (type safety)

**Start with simple enums for things like:**
- Days of week
- Colors
- Sizes (small/medium/large)
- Status (pending/active/completed)
- Directions (north/south/east/west)

As you get comfortable, you can explore more advanced features like bit flags, validation, and JSON support. But for now, just focus on using enums to replace magic numbers and strings in your code!

---

## Basic Enum Pattern

### Simple Enum with iota

```go
type Weekday int

const (
    Sunday Weekday = iota  // 0
    Monday                  // 1
    Tuesday                 // 2
    Wednesday               // 3
    Thursday                // 4
    Friday                  // 5
    Saturday                // 6
)
```

**Key Points:**
- `iota` starts at 0 and increments by 1 for each constant
- Subsequent constants inherit the type and increment automatically
- Creates a custom type (`Weekday`) for type safety

---

## What is `iota`?

`iota` is a predeclared identifier that represents successive untyped integer constants. It resets to 0 whenever the `const` keyword appears and increments after each line.

### iota Reset Behavior

```go
const (
    A = iota  // 0
    B         // 1
    C         // 2
)

const (
    X = iota  // 0 (resets!)
    Y         // 1
    Z         // 2
)
```

---

## Advanced iota Patterns

### 1. Starting from a Different Value

```go
type Month int

const (
    January Month = iota + 1  // 1
    February                   // 2
    March                      // 3
    // ... continues
    December                   // 12
)
```

### 2. Skipping Values

```go
type Status int

const (
    Pending Status = iota  // 0
    _                      // 1 (skipped using blank identifier)
    Active                 // 2
    Inactive               // 3
)
```

### 3. Bit Flags (Powers of 2)

```go
type Permission uint

const (
    Read Permission = 1 << iota  // 1 (binary: 001)
    Write                         // 2 (binary: 010)
    Execute                       // 4 (binary: 100)
)

// Combining flags
const AdminPermission = Read | Write | Execute  // 7 (binary: 111)
```

### 4. Custom Expressions

```go
type Size int64

const (
    KB Size = 1 << (10 * iota)  // 1024
    MB                           // 1048576
    GB                           // 1073741824
    TB                           // 1099511627776
)
```

### 5. Multiple Values on Same Line

```go
const (
    A, B = iota, iota + 1  // 0, 1
    C, D                   // 1, 2
    E, F                   // 2, 3
)
```

---

## String Representation

### Manual String Method

```go
type Color int

const (
    Red Color = iota
    Green
    Blue
)

func (c Color) String() string {
    switch c {
    case Red:
        return "Red"
    case Green:
        return "Green"
    case Blue:
        return "Blue"
    default:
        return "Unknown"
    }
}

// Usage
fmt.Println(Red)  // Output: Red
```

### Using stringer Tool

Go provides a tool called `stringer` to automatically generate `String()` methods:

```bash
go install golang.org/x/tools/cmd/stringer@latest
stringer -type=Color
```

This generates a `color_string.go` file with the `String()` method.

---

## Type Safety

### Without Custom Type (Not Recommended)

```go
const (
    Red = iota
    Green
    Blue
)

func setColor(c int) {
    // Any int is accepted - no type safety!
}

setColor(Red)      // OK
setColor(999)      // Also OK, but wrong!
```

### With Custom Type (Recommended)

```go
type Color int

const (
    Red Color = iota
    Green
    Blue
)

func setColor(c Color) {
    // Only Color type is accepted
}

setColor(Red)           // OK
setColor(999)           // Compile error!
setColor(Color(999))    // Explicit cast required
```

---

## Validation

### Checking Valid Enum Values

```go
type Status int

const (
    Pending Status = iota
    Active
    Completed
    maxStatus  // Sentinel value
)

func (s Status) IsValid() bool {
    return s >= Pending && s < maxStatus
}

// Usage
status := Status(10)
if !status.IsValid() {
    fmt.Println("Invalid status")
}
```

---

## JSON Marshaling

### Basic JSON Support

```go
type Priority int

const (
    Low Priority = iota
    Medium
    High
)

func (p Priority) MarshalJSON() ([]byte, error) {
    return json.Marshal(p.String())
}

func (p *Priority) UnmarshalJSON(data []byte) error {
    var s string
    if err := json.Unmarshal(data, &s); err != nil {
        return err
    }
    
    switch s {
    case "Low":
        *p = Low
    case "Medium":
        *p = Medium
    case "High":
        *p = High
    default:
        return fmt.Errorf("invalid priority: %s", s)
    }
    return nil
}
```

---

## Real-World Example

```go
package main

import "fmt"

// HTTP Status Code Categories
type StatusCategory int

const (
    Informational StatusCategory = iota + 1  // 1xx
    Success                                   // 2xx
    Redirection                               // 3xx
    ClientError                               // 4xx
    ServerError                               // 5xx
)

func (sc StatusCategory) String() string {
    return [...]string{
        "Unknown",
        "Informational",
        "Success",
        "Redirection",
        "Client Error",
        "Server Error",
    }[sc]
}

func (sc StatusCategory) IsError() bool {
    return sc == ClientError || sc == ServerError
}

// Usage
func main() {
    status := ClientError
    fmt.Printf("Status: %s\n", status)
    fmt.Printf("Is error: %v\n", status.IsError())
}
```

---

## Best Practices

1. **Always use a custom type** for enums (not bare `int`)
2. **Start from 0** unless you have a specific reason (easier to detect uninitialized values)
3. **Implement `String()` method** for better debugging and logging
4. **Use bit flags** (`1 << iota`) when you need to combine multiple values
5. **Add validation methods** to check if a value is valid
6. **Document your enums** with comments, especially if values have specific meanings
7. **Use `_` to skip values** you don't want to expose
8. **Keep enums in the same file** as the type they describe

---

## Common Pitfalls

### 1. Forgetting to Create Custom Type

```go
// ❌ Bad
const (
    StatusPending = iota
    StatusActive
)

// ✅ Good
type Status int
const (
    Pending Status = iota
    Active
)
```

### 2. Comparing Different Enum Types

```go
type Color int
type Status int

const Red Color = 0
const Pending Status = 0

// This compiles but is semantically wrong
if Red == Pending {  // Comparing different concepts!
    // ...
}
```

### 3. Not Handling the Zero Value

```go
type Status int

const (
    Active Status = iota + 1  // Start from 1
    Inactive
)

var s Status  // s = 0, which is not a valid Status!
```

**Solution:** Either make 0 meaningful or validate:

```go
const (
    Unknown Status = iota  // 0 is now valid
    Active
    Inactive
)
```

---

## Tools and Libraries

1. **stringer** - Generate String() methods
   ```bash
   go install golang.org/x/tools/cmd/stringer@latest
   ```

2. **enumer** - More powerful enum generator with JSON support
   ```bash
   go install github.com/dmarkham/enumer@latest
   ```

3. **go-enum** - Another popular enum generator
   ```bash
   go install github.com/abice/go-enum@latest
   ```

---

## Summary

- Go uses **const + iota + custom types** for enums
- `iota` auto-increments within a `const` block
- Custom types provide type safety
- Implement `String()` for human-readable output
- Use `1 << iota` for bit flags
- Validate enum values in constructors or setters
- Consider using code generation tools for large enums

---

## Further Reading

- [Go Spec: Iota](https://go.dev/ref/spec#Iota)
- [Effective Go: Constants](https://go.dev/doc/effective_go#constants)
- [Go Blog: Constants](https://go.dev/blog/constants)

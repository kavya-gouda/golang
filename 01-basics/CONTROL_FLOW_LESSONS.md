# Control Flow & Logic - Hands-On Lessons

Runnable, commented Go programs to learn control flow by doing.
Read the theory in [`../02_control_flow/control_flow_guide.md`](../02_control_flow/control_flow_guide.md),
then run these to see it in action and try the "Your Turn" challenges in each file.

## How to run any lesson

Open a terminal in the lesson folder and run:

```bash
go run main.go
```

## Learning path

| # | Folder | What you'll learn |
|---|--------|-------------------|
| 06 | `06_if_else` | `if`, `else`, `else if`, short-statement `if`, logical operators `&& \|\| !` |
| 07 | `07_loops` | The `for` loop in all its forms: classic, while-style, infinite, `range`, `break`/`continue`, labeled break |
| 08 | `08_switch` | `switch` on values, conditionless switch, type switch, `fallthrough` |
| 09 | `09_defer_panic_recover` | Scheduling cleanup with `defer`, raising `panic`, catching with `recover` |
| 10 | `10_error_handling` | The `(value, error)` pattern, sentinel errors, custom error types, `errors.Is` / `errors.As`, wrapping with `%w` |
| 11 | `11_practice` | Worked solutions: FizzBuzz, prime checker, calculator, pattern printing |

## Suggested approach

1. Read the top comment in the `main.go` file for the topic.
2. Run it and match each printed line to the code that produced it.
3. Do the "Your Turn" challenges at the bottom by editing and re-running.
4. When stuck, peek at `11_practice` for complete, working examples.

Tip: change values, break things on purpose, and re-run. Seeing the compiler
error or wrong output is one of the fastest ways to learn.

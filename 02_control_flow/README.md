# Control Flow & Logic in Go

Everything you need to learn — and teach — control flow in Go: a full written
guide, a YouTube tutorial script, and runnable demo programs.

## Contents

| File | Purpose |
|------|---------|
| [`control_flow_guide.md`](control_flow_guide.md) | Complete written reference (concepts, syntax, pitfalls, best practices) |
| [`YOUTUBE_TUTORIAL_SCRIPT.md`](YOUTUBE_TUTORIAL_SCRIPT.md) | Timestamped narration + on-screen actions for recording a video |
| [`demo/`](demo/) | Six runnable programs the tutorial walks through |

## Topics covered

1. **Conditionals** — `if`, `else if`, `else`, short-statement `if`, logical operators
2. **Loops** — the single `for` keyword: classic, while-style, infinite, `range`, `break`/`continue`, labeled break
3. **Switch** — value switch, conditionless switch, type switch, `fallthrough`
4. **Defer / Panic / Recover** — scheduled cleanup and controlled recovery
5. **Error handling** — the `(value, error)` pattern, sentinel errors, custom error types, `errors.Is` / `errors.As`, wrapping with `%w`
6. **Practice** — FizzBuzz, prime checker, calculator, pattern printing

## Run the demos

Each demo is a standalone program. From the `demo/` folder:

```bash
go run 01_if_else.go
go run 02_loops.go
go run 03_switch.go
go run 04_defer_panic_recover.go
go run 05_errors.go
go run 06_practice.go
```

Requires Go 1.18 or newer (tested on Go 1.23).

## For the video

Open `YOUTUBE_TUTORIAL_SCRIPT.md` and follow it section by section. Each demo
file's run command matches a chapter in the script, so you can record straight
through. The script ends with a recording checklist and a ready-to-paste video
description.

## Learn by doing

Read a section of the guide, run the matching demo, then change values and
re-run to see what happens. The fastest way to learn control flow is to break
things on purpose and watch the output shift.

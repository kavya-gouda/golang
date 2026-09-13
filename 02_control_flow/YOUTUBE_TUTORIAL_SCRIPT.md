# YouTube Tutorial Script — Control Flow & Logic in Go

**Target length:** 18–25 minutes
**Level:** Beginner
**Format:** Screen recording with live coding + narration

Legend:
- 🎙️ = what you say (narration)
- 🖥️ = what you show / do on screen
- 💡 = presenter tip (don't say out loud)

---

## 0. Intro (0:00 – 1:00)

🎙️ "Welcome back. In this video we're going to master control flow in Go — that's `if` statements, loops, `switch`, and how Go handles errors and cleanup with `defer`, `panic`, and `recover`. By the end you'll be able to make decisions, repeat work, and handle failure gracefully in any Go program. Let's jump in."

🖥️ Show the folder structure in the editor: `02_control_flow/` with the demo files.

💡 Keep the intro under a minute. Show your face cam if you use one, then cut to the editor.

---

## 1. Why control flow matters (1:00 – 2:00)

🎙️ "Every useful program does three things: it makes decisions, it repeats work, and it deals with the unexpected. Control flow is how we express all three. Go keeps this small and clean — there's just one loop keyword, and errors are plain values, not exceptions. That simplicity is a feature."

🖥️ Nothing to code yet. Optionally show a slide with the 5 topics: if, for, switch, defer/panic/recover, errors.

---

## 2. If / else (2:00 – 5:00)

🖥️ Open `demo/01_if_else.go`. Run it: `go run demo/01_if_else.go`

🎙️ "Let's start with `if`. Notice two things right away: no parentheses around the condition, but the curly braces are always required — even for one line."

🖥️ Point at the basic `if x > 5` block.

🎙️ "We can chain conditions with `else if`. Go checks them top to bottom and runs the first match. Order matters — the highest grade has to be checked first."

🖥️ Point at the grade ladder. Change `score` to 95, re-run, show it prints A.

🎙️ "Here's the pattern you'll use constantly — the short-statement `if`. We declare a variable right in the condition, and it only exists inside this block. This is how Go checks errors."

🖥️ Show `if n := len("hello"); n > 3`. Then show the error-check form in a comment.

🎙️ "And logical operators: `and`, `or`, `not`. Go short-circuits — in `a and b`, if `a` is false, `b` never runs. That's why `if p != nil && p.value > 0` is safe."

💡 Live-edit `age` and `hasLicense` to show the driving rules flip.

---

## 3. Loops with `for` (5:00 – 9:30)

🖥️ Open and run `demo/02_loops.go`.

🎙️ "Go has exactly one loop keyword: `for`. But it wears many hats."

🖥️ Show the classic three-part loop.

🎙️ "Init, condition, post. Standard counting loop. Now watch — if I drop the init and post, `for` becomes a while loop."

🖥️ Show the `for count < 3` version.

🎙️ "Drop the condition entirely and it loops forever. We exit with `break`."

🖥️ Show the infinite loop with `break`.

🎙️ "The one you'll use most is `range`. It walks slices, maps, and strings."

🖥️ Show range over slice, map, string. Run it twice to show the map order changes.

🎙️ "Notice the map printed in a different order the second time. Go randomizes map iteration on purpose — never rely on map order."

🖥️ Show `continue` skipping odds, then labeled `break` exiting nested loops.

🎙️ "A plain break only escapes the inner loop. A label lets you break out of both at once — much cleaner than a boolean flag."

💡 Challenge the viewer: "Pause and write a countdown from 10 to 1."

---

## 4. Switch (9:30 – 13:30)

🖥️ Open and run `demo/03_switch.go`.

🎙️ "`switch` is a cleaner way to write long if-else chains. And here's the big difference from C or Java: Go does NOT fall through. No `break` needed. Once a case matches, we're done."

🖥️ Show the value switch with `case "Saturday", "Sunday"` — point out multiple values per case.

🎙️ "Leave the value off after `switch` and each case becomes a boolean. This reads like a table — great for grade ranges."

🖥️ Show the conditionless switch.

🎙️ "If you actually want the C-style fall-through, you opt in with the `fallthrough` keyword — but it's rare."

🖥️ Show `fallthrough`, run it, explain it prints one then two.

🎙️ "And the type switch — this checks the underlying type of an interface value. Inside each case, the variable already has the right type."

🖥️ Show `describe()` with int, string, bool.

---

## 5. Defer, panic, recover (13:30 – 17:30)

🖥️ Open and run `demo/04_defer_panic_recover.go`.

🎙️ "`defer` schedules a call to run when the function returns — no matter how it returns. Multiple defers run last-in-first-out."

🖥️ Show the LIFO demo output.

🎙️ "The real power is cleanup. Open a file, defer the close right next to it, and you can never forget to close it."

🖥️ Show the cleanup demo.

🎙️ "`panic` stops normal flow, like throwing an exception. Use it only for truly unexpected situations. And `recover`, called inside a deferred function, catches that panic so the program can keep running."

🖥️ Show `safeDivide(10, 0)`. Run it. Point out it recovered and returned -1 instead of crashing.

💡 Optionally delete the recover live to show the crash, then undo.

---

## 6. Error handling (17:30 – 22:00)

🖥️ Open and run `demo/05_errors.go`.

🎙️ "Go doesn't use try/catch for normal errors. Functions return an error value, and we check it. This makes every failure path visible."

🖥️ Show `divide` returning `(float64, error)` and the `if err != nil` check.

🎙️ "For known errors, we define a sentinel — a named error value — and test it with `errors.Is`."

🖥️ Show `ErrInsufficientFunds` and `errors.Is`.

🎙️ "When we need extra data, we make a custom error type — any struct with an `Error() string` method is an error. We pull it back out with `errors.As`."

🖥️ Show `ValidationError` and `errors.As`.

🎙️ "And we can wrap errors with `%w` to add context while keeping the original cause findable."

🖥️ Show the wrapped error and `errors.Is` still matching the root.

---

## 7. Putting it together (22:00 – 24:00)

🖥️ Open and run `demo/06_practice.go`.

🎙️ "Let's combine everything. FizzBuzz uses a loop plus a conditionless switch. The prime checker uses a loop with an early return. The calculator uses a switch plus the error pattern. And the triangle uses nested loops."

🖥️ Run it, scroll through the output.

---

## 8. Outro (24:00 – 25:00)

🎙️ "That's control flow in Go: `if` and `switch` to decide, `for` to repeat, `defer` for cleanup, and error values for failure. All the code is linked in the description. Try the challenges, and if this helped, subscribe for the next one. Thanks for watching."

🖥️ Show the GitHub repo link on screen.

---

## Recording checklist

- [ ] Editor font size bumped to ~18–20pt for readability
- [ ] Terminal font size bumped to match
- [ ] Run each demo once before recording to warm the Go cache (no first-run delay)
- [ ] Close notifications / Do Not Disturb on
- [ ] Zoom level so a full `main.go` fits without horizontal scroll
- [ ] Have the guide (`control_flow_guide.md`) open in a second tab as a reference
- [ ] Mention the repo link verbally and in the description

## Description template (for the video)

```
Learn control flow in Go — if/else, loops, switch, defer/panic/recover, and error handling — with live coding examples you can run yourself.

⏱️ Timestamps
0:00 Intro
2:00 If / else
5:00 Loops (for)
9:30 Switch
13:30 Defer, panic, recover
17:30 Error handling
22:00 Putting it together

💻 Code: https://github.com/kavyagowdam/golang/tree/main/02_control_flow

#golang #programming #tutorial
```

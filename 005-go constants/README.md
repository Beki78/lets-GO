## Constants in Go

### What Are Constants?

Constants in Go are fixed values that cannot be changed during the execution of a program. They are declared using the `const` keyword and are useful for representing values that should remain unchanged throughout the program, such as mathematical constants, configuration settings, or predefined messages. Think of them like those unmovable objects in your life that never change—no matter how hard you try.

---

### Syntax for Declaring Constants

Declaring constants is straightforward. Here’s the basic syntax:

```go
const identifier type = value
```

- `const`: The keyword used to declare a constant.
- `identifier`: The name of the constant.
- `type(optional)`: The data type of the constant.
- `value`: The fixed value assigned to the constant.

### Examples of Constants in Go

Basic Constant Declaration

```go
package main

import "fmt"

func main() {
    const pi float64 = 3.14159
    fmt.Println("The value of pi is:", pi)
}
```

### Multiple Constants

You can declare multiple constants at once, making your code more organized and less painful to read:

```go
package main

import "fmt"

func main() {
    const (
        pi    = 3.14
        e     = 2.71
        earth = "Earth"
    )
    fmt.Println("Constants:", pi, e, earth)
}
```

### Key Characteristics of Constants

1. **Immutable**: Constants cannot be reassigned once declared. You can’t change them, even if you beg and plead.

```go
const a = 10
a = 20 // Compilation error!
```

2. **Compile-Time Evaluation**: Constants are evaluated at compile time, making them faster than variables. So they’re basically the superheroes of your program.

3. **Supported Types**: Constants can hold simple types like:

   - `Booleans` (true, false)
   - `Numbers` (int, float64, etc.)
   - `Strings` ("text")

4. **No Complex Types**: Constants cannot store complex types like slices, maps, or pointers. They must hold simple, immutable values. It’s like a constant commitment to simplicity.

### When to Use Constants?

- For `fixed values` like mathematical constants (π, e).
- For `configuration settings` like timeouts, file paths, or API endpoints.
- For `labels or enumerations` (e.g., days of the week, error codes). You don’t want to confuse your program with random numbers, right?

### Common Pitfalls and Best Practices
#### Pitfalls
1. `Constants Cannot Be Changed`: If you need a value that changes, use a variable instead. Constants don’t care if you change your mind.
2. `Overuse`: Not everything needs to be a constant. Use them for truly fixed values. Don’t turn your code into a constant factory.

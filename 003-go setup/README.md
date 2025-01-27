## Setting Up Go in VS Code

Congratulations! You’re about to enter the world of Go. Before you can start writing Go programs and screaming at the compiler, let’s get your environment ready.

![Let's get ready](https://media3.giphy.com/media/v1.Y2lkPTc5MGI3NjExYmNzbThjeWF4djFobTlhczJka25lbGtnNnZ1djdiYWZmbWJvYTJmYSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/T0aLF2fRB2oR8SFUOV/giphy.gif)

### Step 1: Install VS Code and Extensions

1. Download VS Code
   Get it from https://code.visualstudio.com/.

2. Install Extensions
   Open VS Code and install these extensions:

- **Go Extension Pack**

  - Search for Go by the Go Team in the Extensions Marketplace.
    It handles everything: linting, debugging, IntelliSense, and even helps you feel slightly less dumb while coding.

- **Prettier** - _Code Formatter_
  - Optional, but if you like your code formatted so neatly it could make a robot cry, install this.
- **Error Lens**

  - Optional, Highlights errors and warnings directly in your code, so you can’t pretend they don’t exist.

### Step 2: Initialize Your Go Project

While opening your folder open your terminal and run this command

```go
go mod init my-first-go-project
```

**Explanation**: This command creates a go.mod file to manage your dependencies. Think of it as the project’s birth certificate.

### Step 3: Write Your First Go Program

1. Create a file called main.go.

2. Add the following code:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World! Or should I say, Hello, Go!")
}
```

3. Run your program:

```go
go run main.go
```

Output:

```go
Hello, World! Or should I say, Hello, Go!
```

If it doesn’t work, congratulations—you’ve just met your first Go bug.

### What’s the fmt Package?

The `fmt` package is Go’s way of saying:

“_I’ll handle all your printing needs, but don’t expect me to sugarcoat it._”
Here’s what it can do:

- Print to the Console:

```go
fmt.Println("Go is awesome!")
```

Output: `Go is awesome!`

Formatted Strings:

```go
fmt.Printf("My name is %s and I’ve been coding for %d days.\n", "Beki", 42)
```

Output: `My name is Bereket and I’ve been coding for 42 days.`

- Read Input (if you’re feeling brave):

```go
var name string
fmt.Print("Enter your name: ")
fmt.Scanln(&name)
fmt.Println("Hello,", name)
```

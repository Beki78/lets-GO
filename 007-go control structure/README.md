## Control structure in GO?
Control structures are the tools that help your code make decisions and repeat actions. They’re like the bossy friend who always tells you what to do (but in a good way). Without them, your code would just run in a straight line, like a robot with no sense of direction.

### Types of Control Structures
1.**If/Else Statements** 

* The classic "if this, then that" logic. It’s like deciding whether to wear a jacket based on the weather.

Example:
```go
if temperature < 10 {
    fmt.Println("Wear a jacket, it's cold!")
} else {
    fmt.Println("No jacket needed, you're good!")
}
```
2. **Switch Statements**

* A fancier way to make decisions. It’s like a multiple-choice test for your code.

Example:
```go
switch day {
case "Monday":
    fmt.Println("Ugh, Monday again?")
case "Friday":
    fmt.Println("TGIF!")
default:
    fmt.Println("Just another day...")
}
```
3. **Loops (For)**

* Go’s only loop is the for loop. It’s like a broken record that keeps repeating until you tell it to stop.

Example:
```go
for i := 0; i < 5; i++ {
    fmt.Println("This is loop number", i)
}
```
4. **Defer StatementsDefer Statements**

The "I’ll do it later" of Go. It postpones the execution of a function until the surrounding function returns.

Example:
```go
defer fmt.Println("I’ll run last, even though I’m written first!")
fmt.Println("I’ll run first, even though I’m written second!")
```
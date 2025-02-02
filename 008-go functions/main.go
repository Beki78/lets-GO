package main

import "fmt"

// A simple function that adds two numbers
func add(x int, y int) int {
    return x + y
}

// A function with multiple return values
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return x / y, nil
}

// A function with no return value
func greet(name string) {
    fmt.Printf("Hello, %s! How's it going?\n", name)
}

func main() {
    // Calling the add function
    sum := add(5, 3)
    fmt.Println("5 + 3 =", sum) // Output: 5 + 3 = 8

    // Calling the divide function
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("10 / 2 =", result) // Output: 10 / 2 = 5
    }

    // Calling the greet function
    greet("Gopher") // Output: Hello, Gopher! How's it going?
}
package main

import "fmt"

func main() {
    // Integers
    var age int = 25
    fmt.Println("Age:", age) // Output: Age: 25

    // Floating-point numbers
    var pi float64 = 3.14159
    fmt.Println("Pi:", pi) // Output: Pi: 3.14159

    // Strings
    var name string = "Gopher"
    fmt.Println("Name:", name) // Output: Name: Gopher

    // Booleans
    var isAwesome bool = true
    fmt.Println("Is Go awesome?", isAwesome) // Output: Is Go awesome? true

    // Zero values
    var emptyInt int
    var emptyString string
    var emptyBool bool
    fmt.Println("Empty Int:", emptyInt)       // Output: Empty Int: 0
    fmt.Println("Empty String:", emptyString) // Output: Empty String: 
    fmt.Println("Empty Bool:", emptyBool)     // Output: Empty Bool: false
}
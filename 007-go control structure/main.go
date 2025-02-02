package main

import "fmt"

func main() {
    // If/Else
    temperature := 8
    if temperature < 10 {
        fmt.Println("Wear a jacket, it's cold!")
    } else {
        fmt.Println("No jacket needed, you're good!")
    }

    // Switch
    day := "Friday"
    switch day {
    case "Monday":
        fmt.Println("Ugh, Monday again?")
    case "Friday":
        fmt.Println("TGIF!")
    default:
        fmt.Println("Just another day...")
    }

    // For Loop
    for i := 0; i < 5; i++ {
        fmt.Println("This is loop number", i)
    }

    // Defer
    defer fmt.Println("I’ll run last, even though I’m written first!")
    fmt.Println("I’ll run first, even though I’m written second!")
}
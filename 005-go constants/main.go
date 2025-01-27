package main

import "fmt"

func main() {
    const pi float64 = 3.14159 //single constant
		//multiple constants
		 const (
        g    = 9.18
        e     = 2.71
        earth = "Earth"
    )

    fmt.Println("The value of pi is:", pi)
		fmt.Println("Constants:", g, e, earth)
}
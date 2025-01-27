package main // Defines the main package, which is the starting point of the executable program.

import "fmt" // Imports the fmt package, allowing formatted I/O operations.

func main() { // The main function is the entry point of the program.
	fmt.Println("Hello, World! Or should I say, Hello, Go!") // Prints a greeting message to the console.

	// Prints a formatted message using Printf with a string and an integer.
	fmt.Printf("My name is %s and I’ve been coding for %d days.\n", "Beki", 42)

	var name string // Declares a variable 'name' of type string.
	fmt.Print("Enter your name: ") // Prompts the user to enter their name.
	fmt.Scanln(&name) // Reads the user's input and stores it in the 'name' variable.
	fmt.Println("Hello,", name) // Prints a personalized greeting using the user's name.
}
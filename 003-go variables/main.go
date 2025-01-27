package main

import "fmt"

func main() {
    var mood string = "happy" //here you don't have to explicitly infer to string go is smart he can do it by it self
		
    fmt.Println("Today, I’m feeling:", mood) //prints Today, I’m feeling: happy
		
		health := "bad"
    fmt.Println("Now, I’m feeling:", health) // prints Now, I’m feeling: bad

		var myNumber int
		fmt.Println("How is this" ,myNumber) //prints How is this 0
}

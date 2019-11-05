// My weight loss program.
package main

import "fmt"

// main is the function where it all begins.
func main() {
	var weight float64 = 155
	var age float64 = 36
    fmt.Print("My weight on the surface of Mars is ")
    fmt.Print(weight * 0.3783)
    fmt.Print(" lbs, and I would be ")
    fmt.Print(age * 365 / 687)
    fmt.Println(" years old.")
}

// prints: My weight on the surface of Mars is 58.636500000000005 lbs, and I would be 19.12663755458515 years old.
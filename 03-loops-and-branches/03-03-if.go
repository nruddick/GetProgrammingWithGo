package main

import "fmt"

func main() {
	// assignment:
    var command = "go east"
	// equality:
    if command == "go east" {
         fmt.Println("You head further up the mountain.")
    } else if command == "go inside" {
         fmt.Println("You enter the cave where you live out the rest of your life.")
    } else {
         fmt.Println("Didn't quite get that.")
    }

	var room = "cave"

    if room == "cave" {
        fmt.Println("You find yourself in a dimly lit cavern.")
    } else if room == "entrance" {
        fmt.Println("There is a cavern entrance here and a path to the east.")
    } else if room == "mountain" {
        fmt.Println("There is a cliff here. A path leads west down the mountain.")
    } else {
        fmt.Println("Everything is white.")
    }

}


// Write a guess-the-number program. 
// Make the computer pick random numbers between 1–100 until it 
// guesses your number, which you declare at the top of the program. 
// Display each guess and whether it was too big or too small. 

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const number int = 55

	guessMax := 100
	guessMin := 1

	compGuess := rand.Intn(guessMax)

	for {
		if compGuess == number {
			break
		} else if compGuess > number {
			fmt.Printf("%v is too high\n", compGuess)
			guessMax = compGuess - 1
			compGuess = ((guessMax  + guessMin)/2)	
		} else if compGuess < number {
			fmt.Printf("%v is too low\n", compGuess)
			guessMin = compGuess + 1
			compGuess = ((guessMax + guessMin)/2)
		// } else if compGuess > guessMax {
		// 	fmt.Printf("%v is too high and is not a number between %v and %v\n", compGuess, guessMin, guessMax)
		// } else if compGuess < guessMin {
		// 	fmt.Printf("%v is too low and is not a number between %v and %v\n", compGuess, guessMin, guessMax)
		// } else {
		// 	fmt.Printf("%v is not a number", compGuess)
		}
		time.Sleep(time.Second)
	}
	fmt.Printf("%v was the number!\n", number)
}


// log the time it takes to solve with binary search vs random generator with range limitations
// make testing chart for thousands of tests to compare

















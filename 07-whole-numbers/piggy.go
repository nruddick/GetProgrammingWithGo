/*
days := 365.2425
fmt.Printf("Type %T for %[1]v\n", days)
Experiment: piggy.go
Write a new piggy bank program that uses integers to track the number of cents rather than dollars. Randomly place nickels (5¢), dimes (10¢), and quarters (25¢) into an empty piggy bank until it contains at least $20.
Display the running balance of the piggy bank after each deposit in dollars (for example, $1.05).
*/

package main

import (
    "fmt"
    "math/rand"
)

func main() {
    piggyBank := 0

    for piggyBank < 2000 {
        switch rand.Intn(3) {
        case 0:
            piggyBank += 5
        case 1:
            piggyBank += 10
        case 2:
            piggyBank += 25
        }

        dollars := piggyBank / 100
        cents := piggyBank % 100
        fmt.Printf("$%d.%02d\n", dollars, cents)
    }
}
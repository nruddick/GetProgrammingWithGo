package main

import (
	"fmt"
	"math/rand"
)

// modify to handle leap-years
// generate random year
// for February, assign daysInMonth to 29 for leap years and 28 for other years. You can put an if statement inside of a case block
// use a for loop to generate and display 10 random dates.
var era = "AD"

func main() {
	for i := 1; i <= 10; i++ {
		year := rand.Intn(2020) + 1
		month := rand.Intn(12) + 1
		daysInMonth := 31
	
	switch month {
	case 2:
		if year%4 == 0 {
			daysInMonth = 28
		} else {
			daysInMonth = 29
		}
		fmt.Println(i, era, year, month, day)
	case 4, 6, 9, 11:
		daysInMonth = 30
		day := rand.Intn(daysInMonth) + 1
		fmt.Println(i, era, year, month, day)
	}
	}
}

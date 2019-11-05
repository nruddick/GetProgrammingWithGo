// Experiment: malacandra.go

// Malacandra is much nearer than that: we shall make it in about twenty-eight days.

// C.S. Lewis, Out of the Silent Planet

// Malacandra is another name for Mars in The Space Trilogy by C. S. Lewis. Write a program to determine how fast a ship would need to travel (in km/h) in order to reach Malacandra in 28 days. Assume a distance of 56,000,000 km.

// Compare your solution to the code listing in the appendix.

package main

import (
	"fmt"
)

func main() {
	var distance = 56000000
	var days = 28
	var speed = distance / (days * 24)
	fmt.Println(speed)
}

// 83333
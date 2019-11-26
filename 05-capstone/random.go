package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main () {

	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	fmt.Print(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print(rand.Float64() * 5)
	fmt.Println()

	s1 := rand.NewSource(time.Now().UnixNano())
	fmt.Println(s1)
}
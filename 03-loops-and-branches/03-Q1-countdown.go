package main

import (
    "fmt"
    "time"
	"math/rand"
)

func main() {
    var count = 10

    for count > 0 {
        fmt.Println(count)
        time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			break
		}
		count--
    }

	if count == 0 {
    fmt.Println("Liftoff!")
	} else {
		fmt.Println("Launch failed.")
	}
}
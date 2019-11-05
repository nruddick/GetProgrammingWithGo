// How long does it take to get to Mars?
package main

import "fmt"

func main() {
    const lightSpeed = 299792 // km/s
    var distance = 56000000   // km

    fmt.Println(distance/lightSpeed, "seconds")

    distance = 401000000
    fmt.Println(distance/lightSpeed, "seconds")


	const hoursPerDay = 24
	var speed = 100800      // km/h
	distance = 96300000 // km

	fmt.Println(distance/speed/hoursPerDay, "days")
 }


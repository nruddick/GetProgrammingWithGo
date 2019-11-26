package main

import (
	"fmt"
	"math/rand"
)

func main() {

    // distance in km
    const distanceToMars = 62100000
    // speed in km/s
    const speedMin = 16
    const speedMax = 30
    // a 16 speed price min
    const priceMin = 36
    // a 30 speed price max
    const priceMax = 50

    i := 0

    fmt.Printf("%-17v %-4v %-12v %v \n", "Spaceline", "Days", "Trip type", "Price")
    fmt.Println("==============================================")

    for i < 10 {

        spaceLineId := rand.Intn(3)
        spaceLine := ""
        tripType := "One-way"

        speed := speedMin + rand.Intn(speedMax-speedMin+1)

        // using proportion to calculate price
        price := priceMin + (speed-speedMin)*(priceMax-priceMin)/(speedMax-speedMin)

        // seconds
        tripDuration := distanceToMars / speed
        // days
        tripDuration = tripDuration / (60 * 60 * 24)

        if rand.Intn(2) == 1 {
            tripType = "Round-trip"
            price *= 2
        }

        switch spaceLineId {
        case 0:
            spaceLine = "Space Adventures"
        case 1:
            spaceLine = "SpaceX"
        case 2:
            spaceLine = "Virgin Galactic"
        }

        fmt.Printf("%-17v %4v %-12v $ %4v \n", spaceLine, tripDuration, tripType, price)

        i++
    }
}

// WANTED:
// Spaceline        Days Trip type  Price
// ======================================
// Virgin Galactic    23 Round-trip $  96
// Virgin Galactic    39 One-way    $  37
// SpaceX             31 One-way    $  41
// Space Adventures   22 Round-trip $ 100
// Space Adventures   22 One-way    $  50
// Virgin Galactic    30 Round-trip $  84
// Virgin Galactic    24 Round-trip $  94
// Space Adventures   27 One-way    $  44
// Space Adventures   28 Round-trip $  86
// SpaceX             41 Round-trip $  72

// GOT:
// Spaceline         Days Trip type       Price 
// ==============================================
// Virgin Galactic   25   Round-trip      $   96 
// Virgin Galactic   42   One-way         $   37 
// SpaceX            34   One-way         $   41 
// Space Adventures  23   Round-trip      $  100 
// Space Adventures  23   One-way         $   50 
// Virgin Galactic   32   Round-trip      $   84 
// Virgin Galactic   26   Round-trip      $   94 
// Space Adventures  29   One-way         $   44 
// Space Adventures  31   Round-trip      $   86 
// SpaceX            44   Round-trip      $   72 
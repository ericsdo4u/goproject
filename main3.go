

import (
	"fmt"
)

func main() {
	var totalMiles, totalGallons int
	tripCount := 0

	fmt.Println("Gas Mileage Calculator :)")

	for {
		var miles, gallons int

		fmt.Print("Enter miles driven for trip (or 0 to quit): ")
		fmt.Scanln(&miles)

		if miles == 0 {
			break
		}

		fmt.Print("Enter gallons used for trip (or 0 to quit): ")
		fmt.Scanln(&gallons)

		if gallons == 0 {
			break
		}

		tripMPG := float64(miles) / float64(gallons)
		fmt.Printf("Trip %d MPG: %.2f\n", tripCount+1, tripMPG)

		totalMiles += miles
		totalGallons += gallons
		tripCount++
	}

	if tripCount > 0 {

		averageMPG := float64(totalMiles) / float64(totalGallons)
		fmt.Printf("Combined MPG for all trips: %.2f\n", averageMPG)
	} else {
		fmt.Println("No trips recorded.")
	}
}

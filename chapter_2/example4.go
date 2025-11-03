package main

import (
	"fmt"
	"math"
)

func main() {
	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1+f2+f3
	fmt.Println("The sum is:",sum)

	sum = math.Round(sum*100)/100
	fmt.Println("The sum is now:", sum)

	fmt.Println("Value of Pi is:", math.Pi)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("The circumference is: %.2f",circumference)
}
package main

import (
	"fmt"
	"time"
)

func main() {
	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)

	dayNumber := int(weekday)
	fmt.Printf("The day as a number is %v\n", dayNumber)

	var result string

	switch dayNumber = 0; dayNumber {
	case 1:
		result = "Its Monday"
	case 2:
		result = "Its Tuesday"
	case 3:
		result = "Its Wednesday"
	case 4:
		result = "Its Thursday"
	case 5:
		result = "Its Friday"
	case 6:
		result = "Its Saturday"
	case 7:
		result = "Its Sunday"
	}
	fmt.Println(result)

	x := 0
	switch {
	case x < 0:
		result = "Less than zero"
		// fallthrough
	case x == 0:
		result = "Equals zero"
		// fallthrough
	default:
		result = "Greater than zero"
	}
	fmt.Println(result)
}

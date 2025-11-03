package main

import (
	"fmt"
	"time"
)

func main() {
	exerciseTime := time.Date(2025, time.October, 28, 22, 04, 0, 0, time.UTC)
	fmt.Printf("Doing this exercise on: %s", exerciseTime)

	current := time.Now()
	fmt.Printf("\nCurrent time is %s", current)
	fmt.Printf("\n" + current.Format((time.ANSIC)))
	fmt.Printf("\nThe object's type is %T", current)

	tomorrow := current.AddDate(0, 0, 1)
	fmt.Printf("\n" + tomorrow.Format(time.ANSIC))

	format := "Monday 2006-01-01"
	fmt.Printf("\n" + tomorrow.Format(format))
}

// Write your answer here, and then test your code.
package main

// Uncomment this import section to use required Go packages
import (
	"fmt"
	"strconv"
	// "strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = true
const showHints = true

// calculate() returns the the result of adding 2 numbers
// after parsing them from strings
func calculate(value1 string, value2 string) float64 {
	// Your code goes here.

	// Convert the first string to a float64
	num1, err := strconv.ParseFloat(value1, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
	}
	// Convert the second string to a float64
	num2, err := strconv.ParseFloat(value2, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
	}
	// Calculate and return the result
	answer := num1 + num2
	return answer
}

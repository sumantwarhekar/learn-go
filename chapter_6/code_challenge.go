// Write your answer here, and then test your code.
// Your job is to implement the getCartFromJson() method.

package main

import (
	"encoding/json"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []CartItem {
	var cart []CartItem

	// Create a Reader object from the JSON string
	reader := strings.NewReader(jsonString)

	// Decode the JSON data into the cart slice
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&cart)
	checkError(err)

	return cart
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

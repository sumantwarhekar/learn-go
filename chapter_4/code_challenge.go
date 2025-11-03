// Write your answer here, and then test your code.

package main

import "fmt"

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []CartItem) float64 {
	var sum float64 = 0
	for i, item := range cart {
		fmt.Printf("Item %d: Name: %s, Price: $%.2f, Quantity: %d\n",
			i+1,
			item.Name,
			item.Price,
			item.Quantity,
		)

		sum = sum + (item.Price * float64(item.Quantity))

	}
	return sum
}

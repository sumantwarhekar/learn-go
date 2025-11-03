// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

package main

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// slicesToObjects() returns a slice of Color objects
func slicesToObjects(colorNames []string, hexValues []int) []Color {
	var colors []Color
	n := len(colorNames)
	colors = make([]Color, n)
	for i := 0; i < n; i++ {
		colors[i] = Color{Name: colorNames[i], Hex: hexValues[i]}
	}
	return colors
}

type Color struct {
	Name string
	Hex  int
}

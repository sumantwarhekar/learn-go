package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = make([]string, 0, 3)
	colors = append(colors, "Red", "Green", "Blue")
	fmt.Println("Slice before adding new colors:", colors)
	colors = append(colors, "Purple", "Pink")
	fmt.Println("Slice after adding new colors:", colors)

	colors = remove(colors, 3)
	fmt.Println("Slice after deleting the color on index 3", colors)

	sort.Strings(colors)
	fmt.Println("Slice after sorting", colors)
}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}

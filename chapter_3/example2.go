package main

import "fmt"

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[1])
	fmt.Println("Length of colors array:", len(colors))

	var numbers = [5]int{1, 4, 5, 2, 6, 7}
	fmt.Println(numbers)
	fmt.Println("Length of colors array:", len(numbers))
}

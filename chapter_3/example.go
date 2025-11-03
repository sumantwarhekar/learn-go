package main

import "fmt"

func main() {

	number := 42
	var p *int = &number
	if p == nil {
		fmt.Println("Value of p i nil")
	} else {
		fmt.Println("Value of p is:", *p)
	}
	fmt.Println("Address p is:", p)

	value := 213
	pointer1 := &value
	fmt.Println("\nValue before operation:", *pointer1)
	*pointer1 = *pointer1 / 23
	fmt.Println("Value after operation:", value)
}

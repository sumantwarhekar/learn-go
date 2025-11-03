package main

import "fmt"

func main() {
	i1, i2, i3 := 12, 67, 22
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)

	f1, f2, f3 := 23.43, 34.78, 67.29
	intFloat := f1 + f2 + f3
	fmt.Println("Float sum:", intFloat)

	total := float64(i1) + f2
	fmt.Println("Conversion sum:", total)

	product := float64(i1) * f2
	fmt.Println("Product is:", product)

	num1, num2 := 69, 15
	diff := num1 - num2
	fmt.Println("Differenceis: ", diff)

	num3, num4 := 45, 9
	divison := num3 / num4
	fmt.Println("Divison is:", divison)

	num5, num6 := 98, 17
	remainder := num5 % num6
	fmt.Println("Remainder is:", remainder)

}

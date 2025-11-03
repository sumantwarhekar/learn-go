package main

import "fmt"

func main() {
	var number int = 42
	var str string = "\nThis feels weird\n"

	num2 := 13
	str2 := "\nthis is ok"

	fmt.Print(number, str)
	fmt.Print(num2, str2)

	string1 := "\nThe quick red fox"
	string2 := "jumped over"
	string3 := "the lazy brown dog."
	fmt.Println(string1, string2, string3)

	value := 69
	stringLength, err := fmt.Println("The value is", value)
	if err == nil {
		fmt.Println("String length is", stringLength)
	}

	fmt.Printf("Data type of number is %T", number)
	fmt.Printf("\nData type of string1 is %T", string1)
}

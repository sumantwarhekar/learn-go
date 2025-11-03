package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println("The string given by user is:", str)

	fmt.Print("Enter a number:")
	str2, _ := reader.ReadString('\n')
	num, err := strconv.ParseFloat(strings.TrimSpace(str2), 64)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("The number is:", num)
	}
}

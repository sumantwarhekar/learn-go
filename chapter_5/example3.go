package main

import (
	"fmt"
	"time"
)

func main() {
	go say("Hello from NYC!")
	println("hello from main")
	time.Sleep(3 * time.Second)
	println("All done.")
}

func say(message string) {
	time.Sleep(1 * time.Second)
	fmt.Println(message)
}

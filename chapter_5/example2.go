package main

import "fmt"

func main() {
	dog := Dog{"Poodle", "Woof"}
	dog.Speak()
	dog.Sound = "Bark!"
	dog.Speak()
	dog.SpeakThreeTimes()
}

type Dog struct {
	Breed string
	Sound string
}

func (d Dog) Speak() {
	fmt.Printf("The %v says %v\n", d.Breed, d.Sound)
}

func (d Dog) SpeakThreeTimes() {
	fmt.Printf("%v %v %v\n", d.Sound, d.Sound, d.Sound)
}

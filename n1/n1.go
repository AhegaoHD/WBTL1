package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Speak() {
	fmt.Printf("My name is %s and I am %d years old.\n", h.Name, h.Age)
}

type Action struct {
	Human
	Activity string
}

func (a Action) Do() {
	fmt.Printf("%s is %s.\n", a.Name, a.Activity)
}

func main() {
	person := Action{
		Human:    Human{Name: "John", Age: 30},
		Activity: "running",
	}

	person.Speak() // "My name is John and I am 30 years old."
	person.Do()    // "John is running."
}

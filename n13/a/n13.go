package main

import "fmt"

func main() {
	a := 5
	b := 10

	a = a ^ b
	b = a ^ b // b получит значение a
	a = a ^ b // a получит первоначальное значение b

	fmt.Println(a, b) // Выведет "10 5"

}

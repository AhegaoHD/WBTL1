package main

import "fmt"

func main() {
	a := 5.4
	b := 10.3

	a, b = b, a

	fmt.Println(a, b) // Выведет "10 5"

}

package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	i := 2 // например, удалим 3-й элемент (индекс 2)

	// Удаляем i-ый элемент
	s = append(s[:i], s[i+1:]...)

	fmt.Println(s) // [1 2 4 5]
}

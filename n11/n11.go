package main

import (
	"fmt"
)

// Intersection возвращает пересечение двух множеств
func Intersection(set1, set2 []int) []int {
	set1Map := make(map[int]bool)
	for _, v := range set1 {
		set1Map[v] = true
	}

	var result []int
	for _, v := range set2 {
		if _, exists := set1Map[v]; exists {
			result = append(result, v)
			// Удаляем значение из map, чтобы предотвратить дублирование
			delete(set1Map, v)
		}
	}

	return result
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}
	fmt.Println(Intersection(set1, set2)) // [4 5]
}

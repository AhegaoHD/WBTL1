package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// Разделяем строку на слова
	words := strings.Fields(s)

	// Переворачиваем слова
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Объединяем слова обратно в строку и возвращаем
	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun"
	reversed := reverseWords(s)
	fmt.Println(reversed) // выводит: sun dog snow
}

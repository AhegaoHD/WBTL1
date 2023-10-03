package main

import (
	"fmt"
)

func reverseString(s string) string {
	// Конвертируем строку в слайс рун (элементы Unicode)
	runes := []rune(s)

	// Переворачиваем слайс
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Возвращаем перевернутый слайс как строку
	return string(runes)
}

func main() {
	s := "главрыба"
	reversed := reverseString(s)
	fmt.Println(reversed) // выводит: абырвалг
}

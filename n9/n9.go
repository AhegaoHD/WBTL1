package main

import (
	"fmt"
)

func multiplyByTwo(in <-chan int, out chan<- int) {
	for number := range in {
		out <- number * 2
	}
	close(out)
}

func main() {
	// Инициализация каналов
	input := make(chan int)
	output := make(chan int)

	// Запуск горутины для умножения чисел на 2
	go multiplyByTwo(input, output)

	// Отправка чисел в канал input в отдельной горутине
	go func() {
		numbers := []int{1, 2, 3, 4, 5}
		for _, number := range numbers {
			input <- number
		}
		close(input)
	}()

	// Чтение и вывод результата из канала output
	for result := range output {
		fmt.Println(result)
	}
}

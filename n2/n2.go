package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup // Используем WaitGroup, чтобы дождаться завершения всех горутин
	out := make(chan int, len(numbers))

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done() // Уменьшаем счетчик при завершении горутины
			out <- n * n
		}(num)
	}

	// Запускаем горутину, которая закроет канал out после завершения всех других горутин
	go func() {
		wg.Wait()
		close(out)
	}()

	// Читаем квадраты чисел из канала и выводим их
	for square := range out {
		fmt.Println(square)
	}
}

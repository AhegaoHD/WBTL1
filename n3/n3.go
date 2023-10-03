package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	// Создаем канал для передачи квадратов чисел
	squaresCh := make(chan int, len(numbers))

	// Для каждого числа запускаем горутину, чтобы вычислить его квадрат
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			squaresCh <- n * n
		}(num)
	}

	// Запускаем горутину для закрытия канала после завершения всех горутин
	go func() {
		wg.Wait()
		close(squaresCh)
	}()

	// В основной горутине читаем квадраты чисел из канала и вычисляем их сумму
	sum := 0
	for square := range squaresCh {
		sum += square
	}

	fmt.Printf("Сумма квадратов: %d\n", sum)
}

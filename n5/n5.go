package main

import (
	"fmt"
	"time"
)

const N = 5 // длительность работы программы в секундах

func main() {
	dataCh := make(chan int)

	// Горутина для отправки данных в канал
	go func() {
		counter := 0
		for {
			select {
			case <-time.After(1 * time.Second): // Ждем 1 секунду перед отправкой следующего значения
				dataCh <- counter
				counter++
			}
		}
	}()

	// Горутина для чтения данных из канала
	go func() {
		for val := range dataCh {
			fmt.Println("Прочитано:", val)
		}
	}()

	// Ожидаем N секунд перед завершением программы
	<-time.After(N * time.Second)
	fmt.Println("Завершение программы")
}

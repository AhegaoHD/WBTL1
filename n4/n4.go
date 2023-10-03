package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

const channelSize = 100

func main() {
	rand.Seed(time.Now().UnixNano())

	var numWorkers int
	fmt.Print("Введите количество воркеров: ")
	_, err := fmt.Scanf("%d", &numWorkers)

	if err != nil {
		fmt.Printf("Ошибка при чтении числа воркеров: %s\n", err)
		return
	}

	// Выводим количество воркеров в консоль
	fmt.Printf("Запускаем %d воркеров...\n", numWorkers)

	dataCh := make(chan int, channelSize)

	//Использование контекста позволяет управлять жизненным циклом горутин и завершать их корректно.
	//Когда контекст отменяется (через вызов cancel()), все горутины, которые "слушают" этот контекст, получают сигнал об этом через канал Done() и могут корректно завершить свою работу.
	//Это предоставляет централизованный способ управления жизненным циклом горутин.
	ctx, cancel := context.WithCancel(context.Background())

	// Запускаем горутину для постоянной записи данных в канал
	go produceData(ctx, dataCh)

	// Запускаем воркеров
	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, i, dataCh)
	}

	// Ожидаем сигнал завершения
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh

	// Завершаем контекст, чтобы остановить воркеров
	cancel()
	fmt.Println("Завершение работы...")
	time.Sleep(1 * time.Second) // даем время для корректного завершения горутин
}

func produceData(ctx context.Context, ch chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- rand.Intn(100) // Генерируем произвольное число для демонстрации
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func worker(ctx context.Context, id int, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d завершил работу\n", id)
			return
		case data := <-ch:
			fmt.Printf("Воркер %d прочитал число: %d\n", id, data)
		}
	}
}

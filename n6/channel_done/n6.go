package main

import (
	"fmt"
	"time"
)

func worker(stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			fmt.Println("Worker stopped via channel")
			return
		default:
			fmt.Println("Worker doing work")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ch := make(chan struct{})
	go worker(ch)

	time.Sleep(3 * time.Second)
	close(ch)
	time.Sleep(1 * time.Second)
}

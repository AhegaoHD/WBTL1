package main

import (
	"fmt"
	"time"
)

func worker(timeout <-chan time.Time) {
	for {
		select {
		case <-timeout:
			fmt.Println("Worker stopped due to timeout")
			return
		default:
			fmt.Println("Worker doing work")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	timeout := time.After(3 * time.Second)
	go worker(timeout)

	time.Sleep(4 * time.Second)
}

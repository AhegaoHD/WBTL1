package main

import (
	"fmt"
	"time"
)

func worker() {
	fmt.Println("Worker started")
	time.Sleep(2 * time.Second)
	fmt.Println("Worker returned from function")
	return
}

func main() {
	go worker()
	time.Sleep(4 * time.Second)
}

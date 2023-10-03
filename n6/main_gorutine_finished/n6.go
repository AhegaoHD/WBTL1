package main

import (
	"fmt"
	"time"
)

func worker() {
	for {
		fmt.Println("Worker doing work")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go worker()
	time.Sleep(3 * time.Second)
}

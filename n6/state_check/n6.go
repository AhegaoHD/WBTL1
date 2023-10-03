package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var stop int32

func worker() {
	for {
		if atomic.LoadInt32(&stop) == 1 {
			fmt.Println("Worker stopped via state check")
			return
		}
		fmt.Println("Worker doing work")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go worker()

	time.Sleep(3 * time.Second)
	atomic.StoreInt32(&stop, 1)
	time.Sleep(1 * time.Second)
}

package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[int]int
}

func (sm *SafeMap) Set(key, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Get(key int) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	val, exists := sm.m[key]
	return val, exists
}

func main() {
	smap := &SafeMap{
		m: make(map[int]int),
	}

	var wg sync.WaitGroup

	// Запись данных в map с использованием горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			smap.Set(i, i*i) // записываем квадрат числа
		}(i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		if val, exists := smap.Get(i); exists {
			fmt.Printf("Key: %d, Value: %d\n", i, val)
		}
	}
}

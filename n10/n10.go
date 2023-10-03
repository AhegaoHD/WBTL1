package main

import "fmt"

func main() {
	temperature := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	dict := make(map[int][]float32)
	for _, t := range temperature {
		dict[int(t)/10*10] = append(dict[int(t)/10*10], t)
	}
	for key, temps := range dict {
		fmt.Printf("%d:%v, ", key, temps)
	}
}

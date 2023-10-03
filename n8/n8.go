package main

import (
	"fmt"
)

// Установка i-го бита в 1
func setBit(number *int64, i uint) {
	*number |= (1 << i)
}

// Установка i-го бита в 0
func clearBit(number *int64, i uint) {
	*number &= ^(1 << i)
}

func main() {
	var num int64 = 5 // 101 in binary
	var bitPosition uint = 1

	fmt.Printf("Initial number: %b\n", num)

	setBit(&num, bitPosition)
	fmt.Printf("After setting bit %d: %b\n", bitPosition, num)

	clearBit(&num, bitPosition)
	fmt.Printf("After clearing bit %d: %b\n", bitPosition, num)
}

package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int).Exp(big.NewInt(2), big.NewInt(20+1), nil) // 2^(20+1)
	b := new(big.Int).Exp(big.NewInt(2), big.NewInt(20+2), nil) // 2^(20+2)

	// Сложение
	sum := new(big.Int).Add(a, b)
	fmt.Printf("a + b = %s\n", sum)

	// Вычитание
	sub := new(big.Int).Sub(a, b)
	fmt.Printf("a - b = %s\n", sub)

	// Умножение
	mul := new(big.Int).Mul(a, b)
	fmt.Printf("a * b = %s\n", mul)

	// Деление (целочисленное)
	div := new(big.Int).Div(a, b)
	fmt.Printf("a / b = %s\n", div)
}

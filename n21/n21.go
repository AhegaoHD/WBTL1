// адаптер для сложения, где каждый из аргументов может быть либо строкой, либо числом
package main

import (
	"fmt"
	"strconv"
)

// Целевой интерфейс
type Adder interface {
	Add(a, b interface{}) (int, error)
}

// Адаптируемый класс: сложение двух чисел
type IntAdder struct{}

func (ia *IntAdder) AddInts(a, b int) int {
	return a + b
}

// Адаптер
type UniversalAdderAdapter struct {
	intAdder *IntAdder
}

func (uaa *UniversalAdderAdapter) Add(a, b interface{}) (int, error) {
	intA, err := convertToInt(a)
	if err != nil {
		return 0, err
	}
	intB, err := convertToInt(b)
	if err != nil {
		return 0, err
	}
	return uaa.intAdder.AddInts(intA, intB), nil
}

func convertToInt(value interface{}) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case string:
		return strconv.Atoi(v)
	default:
		return 0, fmt.Errorf("unsupported type")
	}
}

func main() {
	adapter := &UniversalAdderAdapter{
		intAdder: &IntAdder{},
	}

	// Сложение числа и строки через адаптер
	result, err := adapter.Add(5, "7")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result) // 12

	// Сложение двух чисел
	result, err = adapter.Add(3, 4)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result) // 7

	// Сложение двух строк
	result, err = adapter.Add("5", "6")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result) // 11
}

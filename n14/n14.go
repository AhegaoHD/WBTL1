package main

import (
	"fmt"
)

func determineType(i interface{}) string {
	switch v := i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int: // пример с каналом, вы можете расширить для других каналов
		return "channel"
	default:
		return fmt.Sprintf("unexpected type %T", v)
	}
}

func main() {
	variables := []interface{}{
		42,
		"Hello, World!",
		true,
		make(chan int),
		3.14,
	}

	for _, v := range variables {
		fmt.Printf("Value: %v, Type: %s\n", v, determineType(v))
	}
}

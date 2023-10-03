package main

import "strings"

// исходная большая строка (v), созданная функцией createHugeString,
// не будет освобождена сборщиком мусора. Это происходит потому,
// что подстрока (v[:100]) все еще ссылается на исходный массив байтов исходной строки.
var justString string

func createHugeString(size int) string {
	return strings.Repeat("A", size)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
}

func main() {
	someFunc()
}

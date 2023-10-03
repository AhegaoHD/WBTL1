package main

import (
	"fmt"
)

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool)
	for _, s := range strings {
		set[s] = true
	}

	fmt.Println("Множество:")
	for key := range set {
		fmt.Println(key)
	}
}

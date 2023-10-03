package main

import (
	"fmt"
	"strings"
)

func hasUniqueChars(s string) bool {
	seen := make(map[rune]bool)

	for _, char := range s {
		lowerChar := strings.ToLower(string(char))
		if _, exists := seen[rune(lowerChar[0])]; exists {
			return false
		}
		seen[rune(lowerChar[0])] = true
	}
	return true
}

func main() {
	tests := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
	}

	for _, test := range tests {
		fmt.Printf("%s: %v\n", test, hasUniqueChars(test))
	}
}

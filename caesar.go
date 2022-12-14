package main

import (
	"fmt"
	"strings"
)

func cipher(input string, offset int) string {
	var startIndex int
	var wrappedChar rune

	getOffset := func(start, char, shift int) rune {
		return rune(start + (((char - start) + shift) % 26))
	}

	slice := make([]string, len(input))
	for _, c := range input {
		if c < 91 {
			startIndex = 65
		} else {
			startIndex = 97
		}

		if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
			wrappedChar = getOffset(startIndex, int(c), offset)
		} else {
			wrappedChar = rune(c)
		}
		slice = append(slice, string(wrappedChar))
	}

	return strings.Join(slice, "")
}

func main() {
	input := "Hello, citizen."

	fmt.Println(cipher(input, 6))
}

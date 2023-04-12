package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "rubyab"
	s2 := "railsb"

	answer := compareString(s1, s2)
	fmt.Println(answer)
}

func compareString(str1, str2 string) string {
	repetitionChars := ""
	charsMap := make(map[rune]bool)

	for _, ch := range str1 {
		charsMap[ch] = true
	}

	for _, ch := range str2 {
		if charsMap[ch] {
			if !strings.ContainsRune(repetitionChars, ch) {
				repetitionChars += string(ch)
			}
		}
	}

	return repetitionChars
}

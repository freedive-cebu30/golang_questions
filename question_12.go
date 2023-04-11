package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(stringLength("aaaa"))
	fmt.Println(stringLength("鬼滅の刃です"))
}

func stringLength(value string) int {
	return utf8.RuneCountInString(value)
}

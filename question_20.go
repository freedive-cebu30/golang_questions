package main

import (
	"fmt"
	"strings"
)

func main() {
	languages := []string{"ruby", "php", "python", "java"}
	var languages2 []string
	for _, language := range languages {
		languages2 = append(languages2, strings.Title(language))
	}

	fmt.Println(languages2)
}

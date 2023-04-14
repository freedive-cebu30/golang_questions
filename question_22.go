package main

import (
	"fmt"
	"sort"
)

func main() {
	languages := []string{"ruby", "php", "python", "javaScript"}
	sort.Slice(languages, func(i, j int) bool {
		return len(languages[i]) < len(languages[j])
	})
	fmt.Println(languages)
}

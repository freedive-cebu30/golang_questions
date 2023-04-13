package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 12, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5, 4, 3}
	uniqNumbers := make(map[int]bool)

	for _, number := range numbers {
		uniqNumbers[number] = true
	}

	var result []int
	for number := range uniqNumbers {
		result = append(result, number)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	fmt.Println(result)
}

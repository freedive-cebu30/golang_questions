package main

import (
	"fmt"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 8, 9, 10, 18, 20}
	values2 := []int{1, 2, 3, 4, 5, 8, 10, 20}

	answer1 := multiple9(values)
	answer2 := multiple9(values2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}

func multiple9(values []int) int {
	for _, value := range values {
		if value%9 == 0 {
			return value
		}
	}

	return 0
}

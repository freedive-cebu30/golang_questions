package main

import (
	"fmt"
)

func main() {
	values := []int{1, 2, 3, 4, 5}

	doubleValues := make([]int, len(values))
	for i, value := range values {
		doubleValues[i] = value * 2
	}
	fmt.Println(doubleValues)
}

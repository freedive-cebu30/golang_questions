package main

import (
	"fmt"
)

func main() {
	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	maxValue := 0
	
	for _, num := range list1 {
		if maxValue < num {
			maxValue = num
		}
	}
	fmt.Println(maxValue)
}

package main

import "fmt"

func main() {
	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total := 0

	for _, num := range list1 {
		total += num
	}
	fmt.Println(total)
}

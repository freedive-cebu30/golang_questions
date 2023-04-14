package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	values := []string{"1", "12", "12", "13", "2", "3", "4", "5", "6", "7", "8", "9", "10", "5", "4", "3"}

	values2 := []string{"1", "12", "12", "13", "2", "3", "4", "5", "6", "7", "8", "9", "10", "5", "4", "3"}

	sort.Slice(values, func(i, j int) bool {
		n1, _ := strconv.Atoi(values[i])
		n2, _ := strconv.Atoi(values[j])
		return n1 < n2
	})

	fmt.Println(values)

	sort.Strings(values2)
	fmt.Println(values2)
}

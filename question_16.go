package main

import (
	"fmt"
)

func main() {
	list1 := []string{"banana", "apple", "orange"}
	list2 := []string{"mango", "banana", "apple"}
	list3 := []string{"banana", "apple", "orange"}
	list4 := []string{"mango", "banana", "melon"}

	answer1 := compareList(list1, list2)
	answer2 := compareList(list3, list4)
	fmt.Println(answer1)
	fmt.Println(answer2)
}

func compareList(list1, list2 []string) []string {
	map1 := make(map[string]bool)
	for _, v := range list1 {
		map1[v] = true
	}

	map2 := make(map[string]bool)
	for _, v := range list2 {
		map2[v] = true
	}

	var result []string
	for k := range map1 {
		if map2[k] {
			result = append(result, k)
		}
	}

	return result
}

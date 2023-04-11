package main

import "fmt"

func main() {
	showSeason(3)
	showSeason(6)
	showSeason(9)
	showSeason(12)
	showSeason(13)
}

func showSeason(month int) {
	switch month {
	case 1, 2, 12:
		fmt.Println("冬")
	case 3, 4, 5:
		fmt.Println("春")
	case 6, 7, 8:
		fmt.Println("夏")
	case 9, 10, 11:
		fmt.Println("秋")
	default:
		fmt.Println("季節不明")
	}
}

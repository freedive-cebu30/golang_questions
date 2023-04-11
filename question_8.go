package main

import "fmt"

func main() {
	showSeason(3, "japan")
	showSeason(6, "japan")
	showSeason(9, "japan")
	showSeason(12, "japan")
	showSeason(5, "ph")
	showSeason(6, "ph")
}

func showSeason(month int, country string) {
	if country == "japan" {
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
	} else if country == "ph" {
		switch month {
		case 1, 2, 3, 4, 5, 12:
			fmt.Println("雨季")
		case 6, 7, 8, 9, 10, 11:
			fmt.Println("乾季")
		default:
			fmt.Println("季節不明")
		}
	}
}

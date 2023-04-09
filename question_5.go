package main

import "fmt"

func main() {
	var message string
	message = checkAlcohol(19)
	fmt.Printf("あなたの場合は、お酒は%sです\n", message)

	message = checkAlcoholIf(20)
	fmt.Printf("あなたの場合は、お酒は%sです\n", message)
}

func checkAlcohol(age int) string {
	message := "NG"
	if age > 19 {
		message = "OK"
	}
	return message
}

func checkAlcoholIf(age int) string {
	var message string
	if age > 19 {
		message = "OK"
	} else {
		message = "NG"
	}
	return message
}

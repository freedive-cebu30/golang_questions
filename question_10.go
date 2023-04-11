package main

import "fmt"
import "strconv"

func main() {
	fmt.Println(cal(3))
	fmt.Println(cal(4))
}

func cal(number int) int {
	stringNumber := strconv.Itoa(number)
	threeTimesNumber := stringNumber + stringNumber + stringNumber
	intNumber, _ := strconv.Atoi(threeTimesNumber)

	return number + intNumber
}

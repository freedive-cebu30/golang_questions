package main

import "fmt"
import "math"

func main() {
	fmt.Println(calAbs(10, 5))
	fmt.Println(calAbs(10, 13))
}

func calAbs(number1, number2 int) int {
	answer := number1 - number2
	if answer < 0 {
		answer = int(math.Abs(float64(answer)))
	}

	return answer
}

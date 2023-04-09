package main

import (
	"fmt"
	"math"
)

func main() {
	number1 := 10
	number2 := 3
	number3 := 2
	number4 := 10000

	fmt.Printf("%d + %d = %d\n", number1, number2, number1+number2)
	fmt.Printf("%d - %d = %d\n", number1, number2, number1-number2)
	fmt.Printf("%d * %d = %d\n", number1, number2, number1*number2)
	fmt.Printf("%d / %d = %d\n", number4, number1, number4/number1)
	fmt.Printf("%d / %d = %f 余り %d\n", number1, number2, float64(number1)/float64(number2), number1%number2)
	fmt.Printf("%d / %d = %d 余り %d\n", number1, number2, number1/number2, number1%number2)
	fmt.Printf("%dの2乗 = %d\n", number3, int64(math.Pow(float64(number3), 2)))
	fmt.Printf("%dの3乗 = %d\n", number3, int64(math.Pow(float64(number3), 3)))
}

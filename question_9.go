package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(bmi(187.0, 61.0))
	fmt.Println(bmi2(187.0, 61.0))
}

func bmi(height, weight float64) float64 {
	mHeight := math.Pow(height/100, 2)
	return weight / mHeight
}

func bmi2(height, weight float64) float64 {
	return weight * 10000 / math.Pow(height, 2)
}

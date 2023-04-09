package main

import "fmt"

func main() {
	var message string
	message = fizzBuzz(6)
	fmt.Println(message)
	message = fizzBuzz(10)
	fmt.Println(message)
	message = fizzBuzz(15)
	fmt.Println(message)
	message = fizzBuzz(16)
	fmt.Println(message)
}

func fizzBuzz(n int) string {
	var message = "Nothing"

	if n % 15 == 0 {
		message = "FizzBuzz"
	} else if n % 5 == 0 {
		message = "Buzz"
	} else if n % 3 == 0 {
		message = "Fizz"
	}

	return message
}

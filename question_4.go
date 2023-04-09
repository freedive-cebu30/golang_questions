package main

import "fmt"

func main() {
	modulus(5)
	modulus(7)
	modulus(26)
}

func modulus(num int) {
	if num > 25 {
		fmt.Println(num)
	} else {
		for i := 1; i <= 25; i++ {
			if i%num == 0 {
				fmt.Println(i)
			}
		}
	}
}

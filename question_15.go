package main

import "fmt"

func main() {
	showString("world")
}

func showString(str string) {
	fmt.Println(str[0:1])
	fmt.Println(str[len(str)-1:])
	fmt.Println(str[0:2])
	fmt.Println(str[1:4])
}

package main

import "fmt"

func main() {
	fmt.Println("Enter a number :- ")
	var x int
	fmt.Scanln(&x)
	switch x {
	case 10:
		x = 400
		fmt.Println(x)
	case 11:
		x = 100
		fmt.Println(x)
	case 12:
		x = 100
		fmt.Println(x)

	default:
		x = 200
		fmt.Println(x)
	}
}
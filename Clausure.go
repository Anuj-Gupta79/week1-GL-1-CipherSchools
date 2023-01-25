package main

import "fmt"

func main() {
	var Func1 func(int) int
	Func1 = func(x int) int {
		fmt.Println("In a function 1", x)
		return 90
	}
	Q(Func1, 8)

	Func2 := func(x int) int {
		fmt.Println("In a function 2", x)
		return 20
	}(100)
	fmt.Println(Func2)
}

func Q(Func1 func(int) int, u int) {
	Func1(244)
}
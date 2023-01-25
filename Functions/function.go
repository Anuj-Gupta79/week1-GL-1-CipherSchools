package main

import "fmt"

func main() {
	result1, n := Func1()
	fmt.Println(result1, n)

	result2 := Func2(1, 2, 3, 4, 5, 6)
	fmt.Println(result1, n, result2)

	keys := new(int)
	*keys = 21
	var data = 10
	pointer(&data)
	fmt.Println(data) 
}

func Func1() (int, string) {
	return 10, ""
}

func Func2(args ...int) bool {
	for _, value := range args {
		fmt.Println(value)
	}
	return true
}

func pointer(keys *int) {
	*keys = 202
}
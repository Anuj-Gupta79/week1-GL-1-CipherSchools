package main

import "fmt"

func main() {
	var num int
	fmt.Scanf("%d", &num)
	rec(num)
	fmt.Println(num)
}

func rec(num int) {
	if num == 0 {
		return
	}
	if num%2 == 0 { 
		rec(num - 1) 
		fmt.Println(num - 1)
	} else {
		rec(num - 1) 
		fmt.Println(num - 1)
	}
	fmt.Println(num - 1)
}

func fact(num int) int {
	if num == 1 || num == 0 {
		return 1
	}
	if num < 0 {
		return -1
	}
	result := num + fact(num-1)
	return result
}

package main

import "fmt"

func main() {
	arr := []int{10, 5, 99, 100, 34, 90}
	for _, value := range arr{
		fmt.Println(value)
	}
}
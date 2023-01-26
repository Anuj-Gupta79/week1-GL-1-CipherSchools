package main

import "fmt"

func main() {
	var1 := 100
	var2 := &var1
	*var2 = 1000
	fmt.Println(var1, var2)

	var var3 int
	var3 = 100 
	var var4* int
	fmt.Println("X1", var4)

	var4 = &var3
	*var4 = 1000
	fmt.Println(var3, var4)

	var var5 int
	var5 = 100
	var6 := new(int)
	fmt.Println("X2", var6)

	var6 = &var5
	*var6 = 1000
	fmt.Println(var4, var6)

	name := new(string)
	*name = "go"
	fmt.Println(name)
}
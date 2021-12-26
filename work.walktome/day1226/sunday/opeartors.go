package main

import "fmt"

// Golang 运算符

func main() {
	// 1. 算数运算符。
	a := 10
	b := 20
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // r = 2
	fmt.Println(5 % 2) // r = 1

	// 这是单独的语句，不是运算符。
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

}

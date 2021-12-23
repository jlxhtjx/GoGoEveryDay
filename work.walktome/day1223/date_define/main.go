package main

import "fmt"

/*
Today is day 2021-12-23, fighting.
Everything will pass away.
Sometimes I wonder, why can’t I just chase girls directly?
Every time I think: why not.
It's time to study psychology.

 as the people says:智者不入爱河。
*/

func main() {
	fmt.Println("hello 2021-12-23")

	// call print number which is differnet 进制
	printNum()
}

func printNum()  {
	fmt.Println("into printNum function.")
	// a := 10

	var a int = 10
	// 10
	fmt.Print("%d \n", a)
	// 0101
	fmt.Print("%b \n", a)

	// b := 077
	var b int = 077
	fmt.Println("b_o = %o",b)

}
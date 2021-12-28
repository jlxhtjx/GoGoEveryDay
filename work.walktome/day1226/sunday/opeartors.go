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

	// print bit
	printBit()

	// print left move and right move.
	printLeftMove()

	// print value.
	printVlaue()
}

/* Need some time to learn about:
1. Operation system
2. Computer network
3. Data structure
4. Labor Law of the People's Republic of China
5. SQL

*/
func printBit() {
	fmt.Println("----into print bit function.----")
	a := 1             // 001 in 0101.
	b := 5             // 101 in 0101.
	fmt.Println(a & b) // 001 --> 从右往左看，都为1 才是1，so result = 001
	fmt.Println(a | b) // 101 --> 按位与，有一位为1，那么就是1.
	fmt.Println(a ^ b) // 100 --> 不一样，就是1，一样就是0.

}

func printLeftMove() {
	fmt.Println("---left move and right move.----")
	fmt.Println(1 << 2) // 1 --> 100. result = 4
	fmt.Println(4 >> 2) //100 --> 1.  result = 1

}

func printVlaue() {
	fmt.Println("----into print value---")
	a := 5
	a += 5
	fmt.Println(a)
}

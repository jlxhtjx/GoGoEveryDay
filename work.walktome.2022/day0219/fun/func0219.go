package main

import "fmt"

//func.
func main() {
	// func 函数名(参数) (返回值)

	/*
						sayHello()

						//返回参数1
						ret := intSum(1, 2)
						fmt.Println(ret)



					//有参，省略参数类型。
					fmt.Println(intSum2(1, 3))


				ret1 := changeFunc()
				ret2 := changeFunc(10)
				ret3 := changeFunc(10, 20)
				ret4 := changeFunc(10, 20, 30)
				fmt.Println(ret1, ret2, ret3, ret4)


			fmt.Println(funReturnOne())

		// a, b := funReturnTwo()
		fmt.Println(funReturnTwo())
	*/
	//return value with name.
	fmt.Println(funReturnWithName(10, 20))

}

func funReturnWithName(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub

}

//return多个值的时候，用括号括起来。
func funReturnTwo() (int, float32) {
	return 1, 3.14159
}

//return one value.
func funReturnOne() int {
	return 101
}

//可变参要作为最后一个参数。
func changeFunc(x ...int) int {
	fmt.Println("%T", x)
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

//有参有返回，省略参数类型。
func intSum2(x, y int) int {
	return x + y
}

//有参有返回
func intSum(x int, y int) int {
	fmt.Println("--into int sum..--")
	return x + y
}

//无参无返回
func sayHello() {
	fmt.Println("Hello. Sunny boy.")
}

package main

import "fmt"

func main() {
	fmt.Println("hello")
	callManager()

	// 匿名变量不占用命名空间，不会分配内存，所以匿名变量不存在重复声明。
	x, _ := foo()
	_, y := foo()
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
}

// if you go build from the any where, that you must run : go build <GOPATH/>/src/xxx.xx/xxx/xxx

// 定义变量
var heihei string
var name string
var age int
var isMan bool

// 批量定义变量
var (
	a string
	b int
	c bool
	d float32
)

// 变量初始化
var haha string = "哈哈哈"
var myAge int = 24

// 一次初始化多个变量
var herName, herAge = "maomao", 18

// 类型推导
var hisName = "tom"
var hisAge = 22

// 短声明变量：在函数内部，可以使用:= 进行声明并初始化变量
func callManager() {
	n := 100
	m := "mm"
	fmt.Println(m, n)

}

func foo() (int, string) {
	return 10, "ss"

}

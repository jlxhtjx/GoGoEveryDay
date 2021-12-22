package main

import "fmt"


/* 
This is the main function.
Test for var\const\iota.

*/
func main() {
	fmt.Println("hello")
	callManager()

	// if you define a value which haven't use in function. the build will not pass.
	// As for this, must use it in the function or define out of function.
	var unUserValue = 100
	fmt.Println("userValue = ", unUserValue)


	// call unName one 
	unName()

	// print M
	printM()

	// call printLValue
	printValueL()

	// call to print letf move value
	printMB()

	// call muilt iota here.
	// here has a rule: if you want to define value a\b\c must define in function, otherwise can not build pass. Maybe the value a\b\c has been used in golang.
	printMuildIota()
	
}

// if you go build from the any where, that you must run : go build <GOPATH/>/src/xxx.xx/xxx/xxx

// 函数外的变量，只能以关键字开始：var const func ...
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

// 匿名变量
func unName()  {
		// 匿名变量不占用命名空间，不会分配内存，所以匿名变量不存在重复声明。
		x, _ := foo()
		_, y := foo()
		fmt.Println("x = ", x)
		fmt.Println("y = ", y)
}

func foo() (int, string) {
	return 10, "ss"
}



// --------------------------常量 && iota---------------------
// 常量
const pi = 3.1415926
const e = 2.71

// 多常量定义
const (
	pii = 2.222
	kg =58
)

// 统一的常量值
const(
	n1 = 100
	n2
	n3
)

// iota 自增
const (
	m1 = iota
	m2
	m3
	m4
)

func printM()  {
	fmt.Println("m1 = ", m1)
	fmt.Println("m2 = ", m2)

}


//when the value out of func, that it must start with var\const...
const (
	l1 = iota //0
	l2 = 100  //100

	//here is a 插队， even already define l2, but here should be 1, but here is 2. 
	l3 = iota //2
	l4        //3

	// using '_' to ignore some value.
	_         //that means ignore, but iota is still runing, so the l6 is 5 not the 4.
	l6        //5

)

func printValueL()  {
	fmt.Println("l1 = ", l1)
	fmt.Println("l2 = ", l2)
	fmt.Println("l3 = ", l3)
	fmt.Println("l4 = ", l4)
	fmt.Println("l6 = ", l6)

}

const (

	// ignore value = 0
	_ = iota
	
	//here iota = 1, and the left move means 1 --> 1 0*10
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func printMB()  {
	fmt.Println("KB = " , KB)
	fmt.Println("MB = " , MB)
	fmt.Println("GB = " , GB)
	fmt.Println("TB = " , TB)
	fmt.Println("PB = " , PB)
}

// muilt iota define at one line.
const (
	ma, mb = iota + 1, iota + 2    // 1, 2
	mc, md                         // 2, 3
	me, mf                         // 3, 4
)

func printMuildIota()  {


	fmt.Println("ma = %d, mb = %d", ma, mb)
	fmt.Println("mc = %d, md = %d", mc, md)
	fmt.Println("me = %d, mf = %d", me, mf)

}
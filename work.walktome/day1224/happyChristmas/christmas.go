package main

import (
	"fmt"
	"math"
)

func main() {

	// -------------整型 int... -----------
	// 十进制
	var a int = 10
	// 这里如果要打出不同的进制，就得用 Printf，因为这个方法才能打出来不同的进制。
	fmt.Printf("%d \n", a) //10
	fmt.Printf("%b \n", a) //1010 占位符 %b 表示二进制

	// 八进制 以0开头
	var b int = 077
	fmt.Printf("%o \n", b) //77

	// 十六进制，以0x开头
	var c int = 0xff

	// 这里的x和X 应该就是对应 大写和小写。
	fmt.Printf("%x \n", c) //ff
	fmt.Printf("%X \n", c) //FF

	// ---------------浮点型-------
	// float32、float64

	// 在打印浮点数的时候，可以使用fmt的配合动词%f，和c语言一样，可以限制小数点后面的位数等。
	fmt.Printf("%f \n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	// -----------复数------------
	// complex64 complex128
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	// -------------布尔值-------------

	var b1 = false
	if b1 != true {
		fmt.Println("if b1 != true")
		fmt.Println("bool value b1 = ", b1)
	}

	// --------------字符串--------------
	s1 := "hello"
	s2 := "你好"

	// 拼接字符串
	fmt.Println(s1 + " 的中文释义是：" + s2)
	// string length
	fmt.Println("hello的长度 = ", len(s1))

	// -------------byte rune 类型------------
	printByte()

	// --------遍历byte-----
	travesalString()

	// ---------change string-----
	changeString()

	// change the type.
	sqrtDemo()
}

func printByte() {
	/*
		go has two kind of char: byte, rune.
		byte : uint8
		rune : UTF-8  //infact is an int32
	*/
	fmt.Println("------byte------")
	var a = '中'
	var b = 'x'

	// 打印 字符，得用%c
	fmt.Printf("%c \n", a)
	fmt.Printf("%c \n ", b)

}

func travesalString() {
	fmt.Println()
	s := "hello沙河"
	// byte 遍历
	for i := 0; i < len(s); i++ {
		// print 和 printf 是不一样的。
		// 如果用 byte 是打不出 uft8的字符的，所以得用rune
		fmt.Printf("%v(%c) ", s[i], s[i])
	}

	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()

}

func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

// 类型转换
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// I think is look like String.Valueof(xxx * xxx)
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

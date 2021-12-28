package main

import "fmt"

func main() {
	printIf()

	printSimpleIf()

	printFor()

	printForWillBreak()

	printSwitch()
}

func printSwitch() {
	fmt.Println("into function print switch.")
	a := 99
	switch a {
	case 0:
		fmt.Println("a = 0")
	case 1:
		fmt.Println("a = 1")
	case 2:
		fmt.Println("a = 2")
	case 3:
		fmt.Println("a = 3")
	default:
		fmt.Println("unknown")
	}

	// muilt case.
	fmt.Println("into muilt case.")
	// here you can define the var before the sentence. 这里的n 就是判断变量的意思。前面是定义的变量。
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	default:
		fmt.Println("n")
	}

	// switch 判断语句
	fmt.Println("how much age do you?")
	age := 18
	// I found that the var must define the type of int/float/string. It can not mix them. Such as you can not define the string and them ask for >18.
	switch {
	case age > 18:
		fmt.Println("成年")
	case age < 18:
		fmt.Println("未成年")
	default:
		fmt.Println("天山童姥")
	}

	// 1.变量是啥，case 就得是啥，不能混合。不然就编译报错。比如 switch 的变量为 string，然后下面的case 就得是 string类型。
	// var ss string = "aa"

	var ss float32 = 10.0
	// var ss1 int = 18
	// var ss1 float64 = 6.666
	switch ss {
	case 3.14:
		fmt.Println("the switch is aa.")
	/* 2.var ss1 float64 = 6.666
	case ss1: //(float64) 如果 case 的范围比switch 的变量（float32）范围大，然后就编不过去。
	*/

	/* 3.
	var ss int = 18
	case ss:  //switch 变量 //float32.
	 如果在 case 中直接写数值18，会将 int 转成 switch 要求的float32，
	 但是如果直接定义变量为 int 然后case ss，就编不过。
	*/
	case 18:
		fmt.Println("the switch is bb.")
	default:
		fmt.Println("the switch is default.")
	}

}

// the function will stop when the i = 3
func printForWillBreak() {
	fmt.Println("见证奇迹吧。当i = 3的时候，这个函数会停止。")
	for i := 0; i < 10; i++ {
		fmt.Println("奇迹 i = ", i)
		if i == 3 {
			break
		}
	}

	fmt.Println("还是见证奇迹，跳过某一次循环，比如老师点名的时候，看到了请假的同学自动退出。")
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue //当遇到喜欢的女孩子，你的心就会多跳动一次，所以小心哦，不要一下子遇到100个喜欢的女孩子，那么心跳会爆炸的。
		}
		fmt.Println("i = ", i)

	}
}

func printFor() {
	fmt.Println("---into function print For.")
	// for function v1.
	// here is same as if function, which i just using in the for sentence.
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}

	// for function v2.
	// you can just type ";" but the var i must define before for sentence.
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	// for function v3.
	a := 0
	for a < 10 {
		fmt.Println("a = ", a)
		a++
	}

	// for function v4. forever. I think the program will kill the pc, so I just type it but do not run it.
	// but the function can be stop by ("break","goto","return", "panic")

	/*

		for {
			循环语句
		}

	*/

}

//  basic if function.
func printIf() {
	fmt.Println("into print basic if")
	var score = 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

}

// special if wrote.
func printSimpleIf() {
	fmt.Println("into print simple if")
	//what it means here is the var(score) will user only on this if 语句里。
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	fmt.Println("other var score.")
	// watch this score.
	// why here can be same var in a function?
	// 行而不得，反求诸己。从自身寻找原因为什么别人把你排除在外。
	// Ok, because u just a temp var in the last if sentence. So the other can be define and will no collision.
	// In Chinese : 在中国有个同志叫小明，美国也有个同志叫小明，但是中国小明的朋友叫小明，一般只会是中国的小明，美国小明同理。那么有什么办法可以在中国的范围 召唤 美国小明？答案应该是有的。随着我们的学习的递进，应该能逐步学到调取别人的变量。
	var score = 100
	if score >= 90 {
		fmt.Println("A")
	}

}

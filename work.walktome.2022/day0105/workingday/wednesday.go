package main

import "fmt"
// 关于array的 go program.
func main()  {
	
	// Array define.
	defineArray()

	// define the array.
	defineAndSetValue()

	// auto set the array length by the value length. 
	setValueAfterDefine()

	// define an array and set the value by index.
	defineArrayWithIndex()

	//two way to print for array.
	forTheArray()

	// muilt array.
	muiltArray()

	muiltArray2()

	modifyArr()

	//exam
	exam()

	//exam 2
	exam2()
}

func exam2()  {
	//point out the index of the special number.
	fmt.Println("-----into exam 2.------")
	a := []int {1,3,5,7,8}
	var b []int
	b = a
	for i:=0; i<len(a)/2.0;i++{
		for j:=0; j<len(b);j++{
			if ((a[i]+b[j] )==8){
				fmt.Println("sum = 8, index = %d-%d",i,j)
			}
		}
	}
}
// here is test for the array.
func exam()  {
	fmt.Println("----into exam()-----")
	a := []int{1,3,5,7,8}
	var sum int
	for i:=0;i<len(a);i++{
		// fmt.Printf("%d\n",a[i])
		sum = sum+a[i]
	}
	fmt.Println("the result is = ",sum)
}

func modifyArr()  {
	a := [3] int{10,20,30}
	//复制和传参会复制整个数组，所以这里改变之后，只改变副本的值，不改变 原来的值。
	modifyArray(a)
	fmt.Println(a)

	b := [3][2]int {
		{1,1},
		{1,1},
		{1,1},
	}

	//here just modify the b_copy, that the b will not be change .
	modifyArray2(b)
	fmt.Println(b)
}

func modifyArray(x [3]int )  {
	x[0] = 100
	fmt.Println(x[0])
}

func modifyArray2( x[3][2]int)  {
	x[2][0] = 100
}


func muiltArray2()  {

	fmt.Println("-----into muiltArray2----")

	// 第一层可以为空，让编译器推到数组长度，这里就是一个 一维数组，前面的一维数组的长度，可以根据填充的数值来进行定义。第二个数值是在一维数组中，可以放几个string，不是说可以放几个字符的意思。
	a :=[][2] string {
		{"GZ","广府"},
		{"HZ","鹅城"},
		{"SZ","鹏城"},
	}

	//这里的意思应该是定义一个变量 v1, 在a里面的。
	for _,v1 := range a {
		//然后这里的意思应该是在v1 里面取出 v2
		for _,v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}

	// /*
	// 不支持在第二个数组中推到长度。
	b := [3][2]string {
		{"黄焖鸡","18块"},
		{"重庆小面","20块"},
		{"肯德基套餐","34块"},
	}

	fmt.Println("for range.....")

	for _,v1 := range b {
		for _,v2 := range v1 {
						fmt.Printf("%s\t",v2)
			
		}
		fmt.Println()
	}

	// */

}

func muiltArray()  {
	fmt.Println("-----into the function about muilt array.----")
	// here need a ',' in the '{'.
	a := [3][2] string{
		{"BJ","首都"},
		{"NJ","六朝古都"},
		{"SZ","创新城市"},
	}

	// I think is similar to foreach.
	for _,v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}

func forTheArray()  {
	fmt.Println("------------into for function1.---------")
	var a = [] string{"BJ","SH","SZ"}
	for i:=0; i<len(a);i ++{
		fmt.Println(a[i])

	}

	fmt.Println("------------into for function2.---------")
	for index,value := range a{
		fmt.Println(index, value) //这里的index 是新定义的，在 range里面的。
	}
}


func defineArrayWithIndex()  {
	fmt.Println("---------into the function which is set with index.---------")
	a := [] int{1:1, 3:5}
	fmt.Println(a) 
	fmt.Printf("type of a: %T\n",a) 

}

func setValueAfterDefine()  {
	fmt.Println("---------into the function which will set the array length automatically---------")
	var testArray [3] int  //the length of testArray is still 3.But the below numArray is zero, because the numArray haven't set the value.
	var numArray [] int  //here means you can do not define the lenght, and golang will set the length by the value. And same as the below.
	var cityArray = []string {"BJ","NJ","SZ"}
	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Printf("type if numArray : %T\n", numArray)
	fmt.Println(cityArray)
	fmt.Printf("type of cityArray : %T \n", cityArray)
}

func defineAndSetValue()  {
	fmt.Println("---------into define and set value by the default func.---------")
	var testArray [3] int   //array will define the type int and zero value.
	var numArray = [3] int{1,2}  //define an array with pre-value.
	var cityArray = [3] string{"BJ","NJ","SZ"}
	
	// [0 0 0]
	fmt.Println(testArray)
	// [1 2 0]
	fmt.Println(numArray)
	// [BJ NJ SZ]
	fmt.Println(cityArray)



}

func defineArray()  {
	fmt.Println("--------- into the function which is test for array.---------")
	var a [3] int
	// var b [4] int
	//a = b //It can not be done, because the A is different from B.

	//We can use the array as the C program, It can be read from the number, and it's start by 0.

	a[0] = 1001
	a[1] = 1002

	fmt.Println("a[0] = ",a[0])
	fmt.Println("a[1] = ",a[1])
	
}
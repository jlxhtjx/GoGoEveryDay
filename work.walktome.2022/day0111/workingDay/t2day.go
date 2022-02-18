package main

import "fmt"

//slice v2.

/*
和数组的区别就是不需要在中括号不要加数字。

*/

func main() {
	/*  TEST   START
	//索引越界的test
	printSliceIndex()

	//完整的切片表达式
	printFullSlice()

	//make()
	makeSlice()

	//切片不能直接比较
	//切片判空
	chekcSliceNil()



	//cp.
	sliceCopy()



	//遍历数组，其实就是 Java里的List
	sliceRange()


	sliceAppend()


	sliceDoubleCap()

	copySliceByfunc()
		TEST   END*/
	removeFromSlice()
}

func removeFromSlice() {
	//         0   1   2   3   4   5   6
	a := []int{30, 31, 32, 33, 34, 35, 36}

	// 方法：a = append(a[:index], a[index+1:]...)
	//          32
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 33 34 35 36]
}

func copySliceByfunc() {
	fmt.Println("--into slice copy by copy func..--")
	// 如果按以下这种方法，就会将原数组也改变。因为使用的相同的底层数组。
	a := []int{1, 2, 3, 4, 5}
	// b := a
	// fmt.Println(a)
	// fmt.Println(b)
	// b[0] = 1000
	// fmt.Println(a)
	// fmt.Println(b)

	//but if you using the func.copy
	//It will change the slice which is copy, such as slice c.
	c := make([]int, 5, 5)
	copy(c, a)

	fmt.Println(a)
	fmt.Println(c)
	c[0] = 1000
	fmt.Println(a)
	fmt.Println(c)

}

func sliceDoubleCap() {
	fmt.Println("--into print double cap ptr--")
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		//Printf 就是 带格式的输出，这个和 print，println都不一样。
		// %v %d %p 分别表示变量，数值，指针变量。
		//从当前的cap 中可以得出，每次容量(cap) 是成倍增长的。就是 2，4，8，16这样的翻倍增加。
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	var cirySlice []string
	cirySlice = append(cirySlice, "SZ")             //支持追加一个元素。
	cirySlice = append(cirySlice, "GZ", "SH", "BJ") //支持追加多个元素。
	fmt.Println(cirySlice)
	a := []string{"CQ", "HK"}

	cirySlice = append(cirySlice, a...) //支持追加一个切片。
	fmt.Println(cirySlice)
}

func sliceAppend() {
	fmt.Println("-- into slice Append.--")
	var s []int
	fmt.Println(s) //nil now.
	fmt.Println(cap(s))
	s = append(s, 1) //通过 var 声明的零值切片可以直接使用。
	fmt.Println(s)   //[1]
	s = append(s, 2, 3, 4, 5)
	fmt.Println(s)       //[1 2 3 4 5]
	s2 := []int{5, 6, 7} // 这里对比上面 多一个5.
	s = append(s, s2...) // 这里就是把s2 添加到 s后面去, 相当于把一个加了一个切片进去。
	fmt.Println(s)       // [1 2 3 4 (5) 5 6 7]
}

func sliceRange() {
	fmt.Println("--into slice Range.--")
	s := []int{1, 3, 5}
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	//这里的index、 value 都是 从 range 中获得的，应该是抽象了方法，在slice s中 打印index 以及index 对应的value。
	for index, value := range s {
		fmt.Println(index, value)
	}

}

func sliceCopy() {

	fmt.Println("--into slice copy func.--")
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将 s1 赋值给 s1，共用底层数组。
	s2[0] = 100

	fmt.Println(s1) //after s2[0] = 100, It become to [100 0 0], too.
	fmt.Println(s2) //Of course, It should be [100 0 0]
}

func chekcSliceNil() {
	fmt.Print("--into the func to check the --\n")
	//var name []typr
	// 这里只有 var name []T 这种方式的切片是空的，所以如果判空，需要使用 len(s) == 0来判断。
	var s1 []int
	s2 := []int{}
	s3 := make([]int, 0)
	fmt.Println(s1 == nil) // true.
	fmt.Println(s2 == nil) // flase.
	fmt.Println(s3 == nil) // flase.

	fmt.Println("append...")

	s1 = append(s1, 1)
	fmt.Println(s1)
	fmt.Println(s1 == nil)

	s2 = append(s2, 1)
	fmt.Println(s2)
	fmt.Println(s2 == nil)

	s3 = append(s3, 1, 2, 3)
	fmt.Println(s3)
	fmt.Println(s3 == nil) // flase.

}

func makeSlice() {
	fmt.Println("--into the func make.---")
	//create slice 动态ly，we use make().
	// make([]T, size ,cap)
	// T: type of slcie.
	// size : 切片中元素的数量？
	// cap : slice 的容量

	a := make([]int, 2, 10)
	fmt.Print(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	// now we just use for 2/10，容量不会影响当前元素的个数。
	//切片的本质，有点难懂。

}

// nil 就是 golang 里面的 null

func printFullSlice() {
	fmt.Println("---into print full slice.----")
	a := [5]int{1, 2, 3, 4, 5}

	//here need 0 <= low <= high <= max <= cap(a)
	t := a[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
}

func printSliceIndex() {
	fmt.Println("---into print slice Index.---")
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))

	s2 := s[3:4]
	fmt.Printf("s2:%v len(s2):%v cap(s):%v\n", s2, len(s2), cap(s2))
}

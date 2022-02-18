package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)

	var b = []int{3, 7, 5, 3, 6, 7}
	//使用sort 包进行排序
	sort.Ints(b)
	fmt.Println(b)

	//降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(b)))
	fmt.Println(b)
}

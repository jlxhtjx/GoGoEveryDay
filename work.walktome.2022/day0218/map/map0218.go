package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// test for map.
func main() {
	/*  test START
	mapBase()

	//查无此人的判断。
	mapValue()


	//map 的遍历
	mapRange()



	//delete the key value.
	mapDele()



	//按指定顺序遍历
	mapSort()


	//key（元素）为 map 的切片。
	mapSlice1()



	//value（值）为map的切片。
	mapSlice2()



	//统计字符串单词出现的次数。
	calStringWords()

	test END */

	//打印结果
	printResult()

}

func printResult() {
	type Map map[string][]int

	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)         //[1 2 3]
	m["q1mi"] = s                  //这里就是把刚刚的出来的切片s 复制给Map 的key为q1mi 的value.
	s = append(s[:1], s[2:]...)    //这里的... 应该是固定写法  //这里删除了index 为1的切片元素。
	fmt.Printf("%+v\n", s)         //这个应该没有2.  [1 3]
	fmt.Printf("%+v\n", m["q1mi"]) // 为什么 末尾会多一个3？

}

func calStringWords() {
	str := "how do you do"
	strSlice := strings.Split(str, " ") //这是将上面的str 以 " "切分为 string 类型的切片。
	fmt.Println(strSlice)               //[how do you do]
	fmt.Printf("type :%T\n", strSlice)  //type :[]string

	countMap := make(map[string]int, 10) //创建map
	for _, key := range strSlice {       //range 的出来的结果就是 第一个是index，第二个是value.
		//这里把上面拿出来的数组，放到下面作为key，读出里面的value, 判断是否为空。
		_, isReal := countMap[key] //这里是从 map 里面的出来的两个结果，第一个是value，第二个是个bool,判空。
		if !isReal {
			countMap[key] = 1 //这里是如果如果读出来的value为空，那么就设为1.
		} else {
			countMap[key] += 1 // 如果不为空，那么就将里面的value +1.
		}
	}
	fmt.Println(countMap)
}

func mapSlice2() {
	//这里的[]string, 就是一个切片的意思，首先 map[string(这个是key)] (后面跟着的是value) []string
	//这里的map, value 塞了一个 type为 string的[] 切片进去。
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)        //这边是空的。map[]
	fmt.Println(sliceMap == nil) //但是应该和slice一样，使用 make 创建的切片里面不是空的，返回 false.
	fmt.Println("--after init.--")
	key := "China"
	value, ok := sliceMap[key] //这里是设置了一个key，然后使用 value,ok 将value读出来（用sliceMap[key] 返回的值）
	//当然这个key 为 "China" 的map 里值是空的，所以得用 !ok.
	if !ok {
		value = make([]string, 0, 2) //这里用make 创建了一个cap 0 len 2 的切片, 然后塞进一个非空的key的value里。
	}

	value = append(value, "BJ", "SH") //先将里面切片的元素填充。
	sliceMap[key] = value             //对sliceMap[key] 的切片赋值，值为刚刚make出来的切片，并填充了两个元素。这里的key 是上面的China.
	fmt.Println(sliceMap)
}

func mapSlice1() {
	//使用 make 创建一个 元素类型为 map 的slice ： make([]map[string]string, cap), 这里我丽姐 make 里面先以[] 开头就是 slice的意思。
	// 然后 map[后面带中括号带的为key] 后面跟着的为value，在make 中可以定义容量。
	mapSlice := make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	fmt.Println("after init.")
	mapSlice[0] = make(map[string]string, 10) //创建 map
	mapSlice[0]["name"] = "GZ"
	mapSlice[0]["passwd"] = "020"
	mapSlice[0]["address"] = "GDP"

	for index, value := range mapSlice {
		fmt.Printf("index : %d, value: %v\n", index, value)

	}

}

func mapSort() {
	fmt.Println("--into map sord func.--")
	//但是这里为什么要初始化？不是很明白。
	rand.Seed(time.Now().UnixNano()) //rand seed.
	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //%02d 就是 2位整数。
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	//去除map 中所有的key 存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	//对切片进行排序
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

}

func mapDele() {
	//delete map key.
	scoreMap := make(map[string]int, 8)
	scoreMap["SZ"] = 0755
	scoreMap["BJ"] = 010
	scoreMap["GZ"] = 020

	delete(scoreMap, "SZ")
	fmt.Println(scoreMap)
}

func mapRange() {
	//打印 kv 的方法。
	fmt.Println("--print kv.--")
	scoreMap := make(map[string]int, 8)
	scoreMap["SZ"] = 0755
	scoreMap["BJ"] = 010
	scoreMap["GZ"] = 020
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	//If just need k.
	for m := range scoreMap {
		fmt.Println(m)
	}

}

func mapValue() {
	//判断该值是否在这个 map里。
	fmt.Println("--into mapValue.--")
	scoreMap := make(map[string]int, 8)
	scoreMap["SZ"] = 0755
	scoreMap["BJ"] = 010

	//这应该是两个返回值，当这个值存在的时候，ok就会返回 true.
	v, ok := scoreMap["SZ"]
	fmt.Println(scoreMap["SZ"])
	fmt.Println(ok)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

}

func mapBase() {
	fmt.Println("--into mapBase.--")
	//map 应该和Java一样，是使用  key value 的形式。
	//需要使用 make 进行初始化，map 应该在初始化的时候就指定合适的容量。
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"]) // 这个应该是打印"小明" 所在的map 对应的value
	fmt.Printf("type of a : %T \n", scoreMap)

	userInfo := map[string]string{
		"username": "沙河西路",
		"passwd":   "133", //注意这里的值需要加上 ","
	}

	fmt.Println(userInfo)

}

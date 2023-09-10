/*
本节主要讲GO里面的map
同样是一种数据结构，用来表示映射关系
具体可以参考这个博客
https://www.liwenzhou.com/posts/Go/map/
*/

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// map其实和Python里面的字典非常像，是一种无序的、基于key-value的数据结构
// map和前面讲的slice一样是一种引用类型，必须进行初始化才能使用，否则就是个nil（空指针）

func main() {
	map_basic()
	find_keyValue()
	for_range_map()
	delete_function()
	sorted_range()
	complex_map()
	test1()
	test2()
}

func map_basic() {
	a := map[string]int{ // 这里的方法和slice里面一模一样，必须都要使用花括号进行初始化，不然就不能声明一个map变量
		"Tokyo": 2028,
	}
	b := map[string]int{} // 定义一个空的map
	var map_nil_test map[string]int
	fmt.Println(a)
	fmt.Println(b == nil)
	fmt.Println(map_nil_test == nil) // 这里就是个nil

	// 同样的map可以通过make函数进行初始化
	c := make(map[string]int, 10)
	fmt.Println(c)

	// map中添加键值对的操作基本上也和Python字典是一样的
	c["Osaka"] = 2028
	fmt.Println(c)
	fmt.Printf("Variable:%#v\n", c) // v是用来显示变量，前面加上#则是表示更加详细的显示
}

func find_keyValue() {
	a := make(map[string]int, 8)
	a["Tokyo"] = 2028
	a["Osaka"] = 2028
	a["Kyoto"] = 2028
	fmt.Println(a)

	// 判断某个key-value是否在map中
	v, ok := a["china"] // 这里其实可以理解为，对map进行了遍历，如果存在我们想要的键值对，则会返回相应的value和True
	fmt.Println(v, ok)  // 这里显然不存在，所以v返回value(int)的默认值0，ok返回false
}

func for_range_map() {
	// 和slice一样的，可以使用for range来遍历map
	// map是无需的类型，所以遍历出来的顺序和添加的顺序完全无关
	a := make(map[string]int, 8)
	a["Tokyo"] = 2028
	a["Osaka"] = 2028
	a["Kyoto"] = 2028

	for k, v := range a {
		fmt.Println(k, v)
	}
}

func delete_function() {
	// 和slice不一样，map能用delete函数方便的删除元素
	a := make(map[string]int, 8)
	a["Tokyo"] = 2028
	a["Osaka"] = 2028
	a["Kyoto"] = 2028

	delete(a, "Osaka") // 2个参数，map和key

	fmt.Println(a)
}

func sorted_range() {
	// 如果想要按照某种顺序来遍历map呢（因为map本身是无序的，所以普通遍历肯定也是无序）
	a := make(map[string]int, 100)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) // 这玩意就是按照某个特定的格式来生成想要的字符串，感觉上和正则表达式差不多
		value := rand.Intn(100)          // 生成0-99的随机整数
		a[key] = value

	}

	// 按照key从小到大的顺序来遍历map
	keys_in_map := make([]string, 0, 100) // 首先申请一个slice，来存放要进行排序的key（因为slice/数组是存在顺序的）
	for k := range a {
		keys_in_map = append(keys_in_map, k)
	}
	sort.Strings(keys_in_map) // sort包支持各种排序，用其中strings函数对字符串排序

	for _, key := range keys_in_map {
		fmt.Println(key, a[key]) // 其实就是用了slice来当一个中间商
	}
}

func complex_map() {
	// 一些复杂的map类型
	// 元素类型为map的slice
	mapSlice := make([]map[string]int, 8, 100) // 这里注意，只完成了slice的初始化，并没有完成其中map元素的初始化
	// 所以这时候里面的map元素仍然是不可操作的
	fmt.Println(mapSlice[0] == nil) // 还是一个nil
	// 对其中的map元素进行初始化
	mapSlice[0] = make(map[string]int, 8)
	mapSlice[0]["Tokye"] = 2028 // 这样就可以对其中的map元素进行操作了
	fmt.Println(mapSlice[0])

	// value为slice的map
	sliceMap := make(map[string][]int, 8) // 这里和上面的一样的道理，只对外层的map进行了初始化，里面的slice没有
	_, ok := sliceMap["Tokyo"]
	if ok {
		fmt.Println(sliceMap["Tokyo"])
	} else {
		sliceMap["Tokyo"] = make([]int, 2, 10) // 首先对其中的slice元素进行标准化
		sliceMap["Tokyo"][0] = 2023            // 然后在进行slice的相应操作
		sliceMap["Tokyo"][1] = 2028
	}
	fmt.Println(sliceMap["Tokyo"])

}

func test1() {
	// 统计字符串中每个单词出现的次数
	// how do you do
	string_a := "how do you do"

	count := make(map[string]int, len(strings.Split(string_a, " ")))

	for _, word := range strings.Split(string_a, " ") {
		_, ok := count[word]
		if ok {
			count[word] = count[word] + 1
		} else {
			count[word] = 1
		}
	}

	for k, v := range count {
		fmt.Println(k, v)
	}
}

func test2() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	// m["q1mi"] = s
	fmt.Printf("%+v\n", m["q1mi"])
}

/*
理解这里注意，map是一种引用类型，在这里修改过后，s已经变成了[1 3]，按理来说m["q1mi"]输出也应该是[1 3]
但是作为静态语言，m["q1mi"]已经引用了原先s中的三个值，此时s变成只有2个值，并且没有对m["q1mi"]进行重新赋值
所以结果就是，m["q1mi"]中只有前2个值跟随s进行了更新（因为这会s只有2个值了），第三个值仍然保持原来引用的3
如果在输出前对m["q1mi"]进行重新引用，那么结果就是和s一样的[1 3]
*/

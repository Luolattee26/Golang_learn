/*
本节主要讲GO里面的切片
与Python\R不同，GO里面的切片是一个相当灵活的数据类型
具体可以参考这个博客
https://www.liwenzhou.com/posts/Go/slice/
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	array_sum([...]int{1, 2, 3}) // 这样传参数居然也可以hhhhhh
	slice_basic()
	slice_comprare()
	sliceAndCopy()
	slice_explore_all()
	sliceAppend()
	sliceCopyOperations()
	sliceDeleteElements()
	// 两个练习题
	test1()
	test2()
}

func array_sum(x [3]int) { // 从前面的内容可以知道，数组其实不太适合作为一个变量，因为限制太死了
	// 所以如果像这样把数组作为函数的参数的话，其实会受到很多限制（比如长度）
	sum := 0
	for _, value := range x {
		sum += value
	}
	fmt.Println(sum)
}

// 为了解决上述数组的问题，GO里面对array进行了一定的封装，引入了切片
// 实际上，切片的底层还是数组，并且是一个引用类型，包括了引用的内存地址（指针）
func slice_basic() {
	a := []int{}
	b := []string{}
	c := []bool{false, true} // 可以发现slice和array的定义语法基本上一样，也体现了slice的特点——长度不受限制
	fmt.Println(a, b, c)

	// slice本质上还是array，所以也可以通过array来定义切片（实际上也就是切片操作）
	city_go := [...]string{"Tokyo", "Osaka", "Kyoto"}
	city := city_go[0:2]
	fmt.Println(city)
	fmt.Printf("%T\n", city)

	// 也可以针对slice再次进行切片
	city_1st := city[0:1]
	fmt.Println(city_1st)
	fmt.Printf("%T\n", city_1st)

	// 也可通过make函数来创建切片
	// 这里的make函数感觉上有点像var
	country := make([]string, 2, 10) // 第23个参数分别是切片的容量、在内存中分配给这个切片的容量
	// 也就是在这里country虽然是一个长度为2的切片，但是内存里面我们给了它10个长度的容量（append相关）
	// 如果不特殊设置，内存容量与切片容量相等
	country[0] = "Japan"
	country[1] = "US"
	fmt.Println(country)
	fmt.Println(cap(country), "\n", len(country)) // 可以发现len和cap返回的分别是2/10

	// 切片的本质在博客里面有，要注意切片的容量一定是从起始位置算的（如果是利用数组定义的话）
	test1 := [8]int{}
	test2 := test1[0:4]
	test3 := test1[1:4]
	fmt.Println(cap(test2), cap(test3)) // 一个是8，一个是7
}

func slice_comprare() {
	// slice由于是一个引用类型，所以不能直接比较大小/相等，唯一能直接比较的是nil，也就是空指针（引用类型的默认值，大小和容量均为0）
	var a []int // 只声明了，这就是一个nil，如果赋值了哪怕赋的是空值{}，那也不会是nil了
	fmt.Println(a, len(a), cap(a))
	if a == nil {
		fmt.Println("a is nil")
	}
	// 要注意，用make生成出来的，同样不是nil
	// 所以判断slice是不是空的，一般就直接用len来判断
	b := make([]int, 0)
	if b == nil {
		fmt.Println("make function result is nil")
	}
}

func sliceAndCopy() {
	// 和Python中一样，切片本身就是一种浅拷贝操作
	a := make([]int, 3)
	b := a // 当然这里的赋值更是和Python里面的赋值逻辑一模一样
	b[0] = 100
	fmt.Println(a, b)

	c := [...]int{1, 2, 3, 4, 5}
	d := c[0:3]
	d[0] = 10
	fmt.Println(c, d) // 可以发现，对array切片就是一种浅拷贝

}

func slice_explore_all() {
	// 切片的遍历总体上和数组没啥差别
	a := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	for i, v := range a {
		fmt.Println(i, v)
	}
}

func sliceAppend() {
	// 终于到append了，这也呼应了前面为什么要对切片的底层容量进行设置
	// 切片是一个动态类型，当切片的长度处于(cap, infinity)这个区间的时候，会按照一定规则对切片的底层数组进行扩容
	numSlice := []int{}
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)                                                               // append输入的slice对象可以是个nil，不一定需要初始化，返回一个切片
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice) // slice是引用类型，可以直接取出地址
	} // 可以发现，每次超过cap的时候，就会把cap变为2倍，并且可以发现内存地址也发生了改变，证明底层数组发生了变化

	a := []int{213, 234134, 123}
	a = append(numSlice, a...) // 如果是用append实现slice相连，记得加上三个点，这是固定语法
	fmt.Println(a)

	// GO语言的切片扩容规则挺复杂的，这保证了GO的性能
	// 具体的规则可以在这里面看源代码
	// $GOROOT/src/runtime/slice.go
	// 动态扩容是会损耗性能的，所以要在定义的时候就尽量提前订好需要多少的最大容量
}

func sliceCopyOperations() {
	// 这里是slice的copy操作，其实和Python的deep_copy一样
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 10)
	b := a
	copy(c, a)     // a to c
	fmt.Println(a) // [1 2 3 4 5]
	fmt.Println(c) // [1 2 3 4 5]
	c[0] = 1000
	fmt.Println(a) // [1 2 3 4 5]
	fmt.Println(c) // [1000 2 3 4 5]
	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", c) // 两个的内存地址也是不一样的
	fmt.Printf("%p\n", b) // 直接赋值那就是浅层拷贝，指针地址一样
}

func sliceDeleteElements() {
	// GO里面没有提供方便快捷的删除元素的方法......一般通过copy实现
	a := []string{"Tokyo", "Osaka", "Kyoto", "Yokohama"}
	a = append(a[:2], a[3:]...)
	fmt.Println(a)

	// 公式就是要删除第index个元素
	// a = append(a[:i], a[i+1:]...)
}

func test1() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}

func test2() {
	a := [...]int{3, 7, 8, 9, 1}
	sort.Ints(a[:]) // 没有返回值，直接用就可以，注意一下用法，可以设定对哪个范围进行排序
	fmt.Println(a)
}

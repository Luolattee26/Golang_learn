/*
终于来了，数组
具体可以参考这个博客
https://www.liwenzhou.com/posts/Go/array/
*/

// go里面的数组是同一类型数据的集合，定义好后能修改元素，但是大小是不能变的
// 定义的语法也比较直观，两个变量，大小+数据类型

package main

import "fmt"

var a_array [3]int // 这两个数组是不同的，因为大小不一样
var b_array [2]int // 数组的类型是包括大小+数据类型两部分的

func main() {
	fmt.Println(a_array, b_array)
	array_init()
	array_iter()
	matrix()
	array_character()
}

func array_init() {
	// 和普通变量一样的，数组也需要进行初始化，不然就是0之类的默认值
	city_array := [4]string{"北京", "上海", "深圳"} // 这个定义值的方法挺特殊的，注意后面用的是花括号，是一个列表
	fmt.Println(city_array)                   // 如果我们赋值的内容和数组长度对不上，那会自动从前到后赋值初始化

	// 上面那种方法每次都要数一下到底多少个元素，也可以让编译器自己推导有多少个元素
	boolArray := [...]bool{true, false, true, false, true, false, true, false, true, false}
	fmt.Println(boolArray)
	fmt.Println(boolArray[1]) // index，[a:b]切片操作和Python一样同样是只取下界不取上界

	// 也可以使用数组索引的方式来给数组进行初始化
	langArray := [...]string{1: "Golang", 3: "python", 5: "java"} // 长得好像字典......
	fmt.Println(langArray)                                        // 字符串默认值是空字符串，所以显示效果可能不大好
	fmt.Printf("%T\n", langArray)
}

func array_iter() {
	// 数组的遍历首先可以用最基础的for循环来写
	city_array := [...]string{"Tokyo", "Osaka", "Kyoto"}
	for i := 0; i < len(city_array); i++ { // 这个len其实和Python完全一样的
		fmt.Println(city_array[i])
	}

	// 也可以go提供的一个for range的方式
	// 这个方式说实话和Python里面的enumerate一模一样，只不过这里的range被做成了一个保留字
	// 只想要value就可以用  _, value
	for index, value := range city_array {
		fmt.Println(index, value)
	}
}

func matrix() {
	// go原生支持多维数组，这不就是矩阵吗
	city_array := [2][3]string{ // 定义语法和np.array基本上一模一样
		{"Tokyo", "Kyoto", "Osaka"},
		{"Hangzhou", "Nanking", "Shanghai"}, // 多维数组最后的花括号必须连在一起，或者是跟一个逗号做分割
	}
	fmt.Println(city_array)       // 打印出来其实就是一个嵌套
	fmt.Println(city_array[0][0]) // Tokyo

	// 二维数组的遍历
	for _, v1 := range city_array {
		for _, v2 := range v1 { // 因为只有value才是我们的数组，所以这里range的对象是v1
			fmt.Println(v2)
		}
	}

	city_array_dot := [...][3]string{ // 定义的时候要注意只有第一维数组才可以用推导的方法来写
		{"Tokyo", "Kyoto", "Osaka"},
		{"Hangzhou", "Nanking", "Shanghai"},
	}
	fmt.Println(city_array_dot)
}

func array_character() {
	// go语言里面的数组是一种“值类型”
	x := [...]int{1, 2, 3}
	fmt.Println(x)
	foo(x)
	fmt.Println(x)
	// 可以发现前后的x打印出来的并没有发生改变
	// 其实值类型很好理解，就是Python中的deepcopy
	// 不管是将x赋值给另一个数组或者是当成参数
	// 都是直接完整copy一份拿去用，内存里面存储的x的内容都不会发生任何改变
}

func foo(a [3]int) { // 方便演示用了一个函数，具体内容后面会有
	a[0] = 100
}

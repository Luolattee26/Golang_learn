/*
本节主要讲GO里面的函数结构体的基础
具体可以参考这个博客
https://www.liwenzhou.com/posts/Go/function/
*/

package main

import (
	"fmt"
)

// Go语言可以说主体就是各种函数，函数可以作为任何东西出现
// Go语言的函数不存在默认参数这么个东西

var global int = 2028

func main() {
	fmt.Println(intSum(4, 6))
	fmt.Println(changeable_arg(10, 30)) // 可以发现这样就可以随意传入参数，传入的多个参数自动汇总到我们设置好的a这个slice里面
	fmt.Println(changeable_arg(10))
	fmt.Println(changeable_arg(10, 20))
	fmt.Println(changeable_arg(10, 20, 30))
	fmt.Println(return_value(10, 5))
	fmt.Println(defer_function(7, 5))
	variable_()
	fmt.Println(global)
	a := variable_        // 函数可以作为任何东西，比如这里作为了一个变量赋值给了a，注意不要有调用的括号
	fmt.Printf("%T\n", a) // 是func类型
	a()
	fmt.Println(calc(166, 200, add))
}

func intSum(x int, y int) (res int) { // 定义的方式和别的编程语言基本上差不多，但是要注意静态语言的特点
	// 这里括号外的int，其实是定义了返回值的类型（也可以有名字），类型设置错了，就会报错
	res = x + y // 如果上面给返回值定义了名字，那函数体里面就可以直接使用返回值而不需要进行声明了，下面也可以直接用return
	return
}

func changeable_arg(b int, a ...int) (res int) { // 在参数后面跟上...，就是表示可变参数的意思
	// 如果存在可变参数与固定参数混合，那固定参数一定要在可变参数前面
	fmt.Printf("%T\n", a) // 可以发现，go的可变参数实际上就是吧参数做成了一个slice，类型为我们设置好的int
	res = b
	for _, v := range a {
		res = res + v
	}
	return
}

func return_value(a, b int) (sum int, sub int) { // 在这里，我们对参数ab的类型进行了简写，并且设置了多返回值
	sum = a + b
	sub = a - b
	return
}

func defer_function(a, b int) (res int) { // go的函数里面可以使用defer语句来延迟执行，并且按照逆序延迟执行
	res = a + b
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	return // 并且difer语句是在return触发前，执行，或者是最后一条语句触发之后（总而言之就是函数结束前）
}

func variable_() {
	var global int = 1000
	fmt.Println(global) // 函数外用var定义全局变量，函数内用:=/var定义函数内的变量
	// go里面的变量逻辑是从小到大的，先找局部，局部没有就用全局变量
	// 函数里面定义的局部变量，也是不能够在外部进行访问的

	for i := 0; i < 5; i++ {
		fmt.Println(i) // 像这里，i也只在for循环里面生效，哪怕到了函数中都是不起作用的
	}
}

func add(x, y int) int {
	return x + y
}

func calc(x, y int, op func(int, int) int) int {
	// 这里其实并不难理解，我们将函数作为参数op传入了calc这个函数中
	// 这里指定的类型，就是函数所属的func类型，并且注意由于op同样是函数，所以也要制定参数、返回值类型为int
	// 在上面的fmt.Println(calc(166, 200, add))中，我们就把add这个函数传入了op参数中
	// 注意add和op参数的类型要一样，不然肯定是不可能传入的
	// 最终实现了一个联合计算
	return op(x, y)
}

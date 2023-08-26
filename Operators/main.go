/*
本节主要讲各种运算符号
具体可以参考这个博客
https://www.liwenzhou.com/posts/Go/operators/
*/

package main

import "fmt"

func main() {
	calc_operator()
	relation()
	logistic()
	position()
	assign_value()

}

func calc_operator() {
	// 算数运算符，基本上和Python一模一样
	a := 20
	b := 5
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ // 累加，算是一个单独的语句，但是每次是加1，也有--
	fmt.Println(a)
}

func relation() {
	// 关系运算符，也基本上一样
	fmt.Println(10 == 2)
	fmt.Println(10 != 2)
	fmt.Println(10 >= 2)
}

func logistic() {
	// 逻辑运算符，也基本上一样
	fmt.Println(5 > 4 && 10 == 5)
	fmt.Println(!(5 >= 4))
	fmt.Println(10 <= 2 || 5 >= 4)
}

func position() {
	// 位运算符，这个比较抽象的新概念
	// 任何数据在内存里面都是二进制储存的
	// 位运算符就是针对二进制来进行判断的
	a := 1             // 二进制001
	b := 5             // 二进制101，这里是为了方便比较，对齐了一下
	fmt.Println(a & b) // 对比对象的二进制表示，从右往左开始对比，如果都是1，则返回1，不都是，则返回0
	// 所以这里的输出其实应该是001的，和位数更多的二进制齐平，只不过这里001省略为了1
	fmt.Println(a | b) // 这里输出应该是101，但是Println默认十进制输出，所以变成5
	fmt.Printf("%b\n", a|b)
	fmt.Printf("%b\n", a^b)  // 这个其实就是not，只有不一样才输出1，应改为100
	fmt.Printf("%b\n", a<<2) // 在二进制下左移2位，变成100
	fmt.Println(a << 2)      // 变成100转化为10进制的4
	fmt.Printf("%b\n", 4>>2) // 在二进制下右移2位，变成1
	fmt.Println(4 >> 2)      // 还是1
}

func assign_value() {
	// 赋值运算符和Python里面也基本一样
	// 只不过增加了位运算的内容
	a := 5
	a += 5
	fmt.Println(a)
	// 位运算赋值的内容其实基本上是一样的，就是先进行位运算而后赋值，一样的去理解过程就好了

}

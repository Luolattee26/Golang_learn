/*
本节主要讲各种变量的类型
具体可以参考这个博客
https://www.liwenzhou.com/posts/Go/datatype/
*/

package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// 十进制打印成二进制（这个用fmt包来实现，其实很像Python里面一些列操作）
	n := 10               // 之前提到的定义短变量来临时使用
	fmt.Printf("%d\n", n) // 十进制%d
	fmt.Printf("%b\n", n) // 二进制%b

	a := 075 // 八进制可以直接0开头来这样定义
	fmt.Printf("%d\n", a)
	fmt.Printf("%o\n", a) // 八进制%o，上面075中0只是一个定义符号，打出来还是75

	f := 0xff // 八进制可以直接0开头来这样定义
	fmt.Printf("%d\n", f)
	fmt.Printf("%x\n", f) // 十六进制%x
	fmt.Println(f)        // Println默认10进制

	// 比如说定义一个8位二进制的无符号整型变量
	var age uint8
	// age = 1000，这样赋值就会错，因为超过255了（8位二进制）
	fmt.Println(age)

	// 如果像之前那样简单的赋值一个int类型，则会根据运行系统自动选择32/64位
	var test int = 12
	fmt.Printf("%T\n", test) // 这是利用Printf和占位符查看变量类型的语句
	fmt.Println(math.MaxInt)
	fmt.Println(math.MaxInt64) // 也可以发现这里打印出来的两个max是一模一样的

	// 浮点数
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64) // math包有一系列数学相关的操作

	foo()
	multi_lines()
	stirng_operation()
	runeOrbyte()
}

func foo() {
	// 复数
	var com1 complex64  // 实部和虚部都是float32
	var com2 complex128 // 和上面类似
	com1 = 0.1212 + 0.12312i
	com1 = 12121226 + 12121226i
	fmt.Println(com1, com2)
}

func multi_lines() {
	s1 := `
	两个反引号中间的就是多行字符串
	键盘最左边的键，不是单引号！
	里面的东西一律原样输出，转义符\是没用的
	`
	fmt.Println(s1)
}

func stirng_operation() {
	s1 := "tokyo"
	s2 := "i want to go to"
	fmt.Println(s2 + " " + s1)         // 拼接字符串也比较简单
	s3 := fmt.Sprintf("%s %s", s1, s2) // Sprintf这函数能通过特定格式来格式化字符串
	fmt.Println(s3)

	// 字符串的切割
	s4 := "i will go to Tokyo or Osaka"
	fmt.Println(strings.Split(s4, " ")) // strings包自带的函数，和Python基本一样
	fmt.Printf("%T\n", strings.Split(s4, " "))
	// 可以发现返回的是一个[]string类型，这是字符串切片数据类型，后面会涉及到

	// 判断是否包含
	fmt.Println(strings.Contains(s4, "Kyoto"))

	// 判断是否有前缀后缀（啥开始啥结束）
	fmt.Println(strings.HasPrefix(s4, "i"))
	fmt.Println(strings.HasSuffix(s4, "Osaka"))

	// 判断子串出现的位置
	fmt.Println(strings.Index(s4, "Tokyo"))
	fmt.Println(strings.LastIndex(s4, "Tokyo")) // 最后一次出现

	// join操作，字符串切片的连接操作
	// 这是针对字符串切片类型的，和前面的字符串连接用+不一样
	s5 := []string{"i", "Tokyo", "Osaka", "Kyoto"}
	// 字符串切片这个数据类型后面会讲的
	fmt.Println(s5)
	fmt.Println(strings.Join(s5, "-"))
}

func runeOrbyte() {
	// 组成字符串的每个元素都是字符
	// 可以用''单引号来定义
	// byte uint8的别名 表示ASCII码（英文和数字）
	// rune int32的别名 表示UTF-8（中日韩）

	var b byte = 'a' // 英文就1个
	var c rune = '京' // 一个中文字符就占3个字节，所以后面遍历中英文混杂字符串，需要特殊手段

	fmt.Println(b, c) // 打出来的实际上就是ASCII/UTF-8的编码
	fmt.Printf("%T %T", b, c)
	// 可以发现实际上就是uint8 int32的数据类型，只是起了一个别名，来提升代码可读性

}

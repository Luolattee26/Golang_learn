/*
本节主要讲很核心的内容，各种循环、流程
具体可以参考这个博客
https://www.liwenzhou.com/posts/Go/control-flow/
*/

package main

import "fmt"

// go语言里面if的写法基本上和Python一样，只是格式要求较高

func main() {
	if_control_basic()
	if_advance()
	for_circle()
	for_acvance()
	switch_case()
	switch_advance()
	switchDemo5()
	gotoDemo1()
	gotoDemo2()
}

func if_control_basic() {
	score := 65
	// go的自动补全太牛逼了吧
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func if_advance() {
	// 这里其实就是在if语句块内部定义了一个score变量，这样就可以让代码更加简洁有条理
	// 和函数内部的短变量相同的性质，这个变量只会存在于if语句块中，甚至不会到函数中
	// 注意要用分号来分开两个不同的部分（变量和状况），这个和Python不同
	if score := 77; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func for_circle() {
	// go语言里面循环有且仅有for一种，不像其他的存在while啥的
	// for循环的写法其实和if的大同小异
	// 可以注意下面这种for循环的写法，是go语言中最基础的写法
	// 条件语句分为三部分，初始值+循环条件+迭代
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 10
	for ; i < 16; i++ {
		// 初始的赋值语句可以省略掉的，但是那个分号必须要保留的
		fmt.Println(i)
	}

	// 在go里面也可以只写一个结束条件语句，这样其实就和别的语言的while差不多了
	a := 5
	for a > 0 { // 当只有一个条件语句的时候，就可以不写分号了，注意一定是只有一个条件语句的时候
		// 所以这里a--需要放在循环体里面
		fmt.Println(a)
		a--
	}
}

func for_acvance() {
	// 无限循环，直接不写任何条件就行
	// for {

	// }

	// go里面同样存在break、continue操作
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i >= 3 {
			break // 注意break、continue一定是针对循环的，这是和Python一样的
		}
	}

	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // 一定要注意，continue的作用是跳过本次循环
		}
		fmt.Println(i)
	}

}

// go里面为了防止出现大量if的情况，还提供了switch用来做多情况选择
func switch_case() {
	finger := 7
	// 整体语法的缩进有点奇怪，但是可读性还是OK
	switch finger { // 这里是要进行判断的表达式
	case 1: // case后面直接跟对应的情况的函数值就可以了
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default: // 每个switch结构体都只允许有一个default
		fmt.Println("无效的输入！")
	}

}

func switch_advance() {

	switch n := 9; n + 1 { // 前面提到的循环语法也是完全适用于switch的
	// 当然这个变量表达式可以不放在switch里面，可以放在case那也行的，只不过要提前定义一个变量了
	case 8, 9: // 可以同时存在多个判断值
		fmt.Println("yes")
	case 10:
		fmt.Println("no")
	}
}

// fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
func switchDemo5() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

/*
goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。
Go语言中使用goto语句能简化一些代码的实现过程。 例如双层嵌套的for循环要退出时：
*/
func gotoDemo1() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			break
		}
	}
}

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag // 触发后就会直接退出整个循环结构（可以发现外面的大循环也退掉了）
				// 如果把标签写在大循环里面的话，就只会推出小循环了，逻辑是根据标签所在的层级来判断的
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return // 这个return在这个例子里面其实没有实际作用的
	// 标签
breakTag: // 标签的话是一种特殊的语句，通过顶格写来定义
	fmt.Println("结束for循环")
}

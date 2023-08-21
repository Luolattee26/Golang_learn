package main

import "fmt"

// import 也最好放在最前面

// go里面的变量必须要经过声明才能用，不能像Python那样，而且必须要定义好类型，像vb~
// 同一个[作用域]内部不允许重复申明同样名字的变量

var (
	a int = 2028 // 申明的同时设置好初始值
	b string
	c = true // go语言编译器支持类型推导，所以直接赋值也是可以的，不一定要先设置类型
	d float32
) // 变量批量申明，和import是基本一样的操作啦
// 申明之后会默认在内存里面对其进行初始化，先赋值一个默认值，bool的默认是false，[指针]是nil

func main() {
	b = "For Tokyo"
	d = 2000.1226
	fmt.Println(a, b, c, d)

	// 在函数内部，我们也可以使用“短变量”，短变量是仅限于函数体内部的变量，走不出去
	m := "For Osaka"
	fmt.Println(m)

	foo()

	/*
		go里面也存在匿名变量 _
		当我们只需要一个变量来占位置的时候，比如函数返回两个只只需要其中一个取出来
		就可以使用下划线来占位
		匿名变量不占用命名空间+不使用内存，所以直接用就可以了
	*/
} // 其实可以发现，当程序run的时候，写在main函数里面的东西会自动运行，这和Python是不一样的，也是init/main这两个函数的特殊性

/*
可以发现，在go里面但凡是函数外面的语句，都只能用关键字开头，比如var、const、import、package
不能像Python一样整一个x=100随便到处放
*/

func foo() {
	const (
		pi  = 3.14
		e   = 2.71
		e_1 // 没有赋值，那就默认和上面的一个值一样
	)

	const (
		n1 = iota // iota这是个关键字，iota[碰到一个const就会重置一次成0]，在该const里面每增加一行，iota就会加1
		_         // iota常常用来实现枚举、计数等等，也会结合匿名变量来使用
		n2
		n3 = 100
		n4 = iota // 这里不写iota就会变成100
	)

	const n5 = iota

	fmt.Println("Tokyo Osaka Kyoto")

	fmt.Println(pi, e, e_1)

	fmt.Println(n1, n2, n3, n4, n5) // 正确输出应该是0，2,100,4，0
	// 这就是定义常量，因为是常量，所以在申明的时候就必须赋值，这个要注意

	/*
		一行内定义多个常量的情况
	*/
	const (
		i1, i2 = iota, iota + 1
		i3, i4
		i5, i6
	)
	fmt.Println(i1, i2, i3, i4, i5, i6) // correct = 0, 1, 1, 2, 2, 3
}

// 这里看foo()里面两个打印就能发现，main自动运行了，带着里面的foo也跑了一次，但是foo是不会自动运行的，所以整体也就打印了一次

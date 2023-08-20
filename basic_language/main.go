/*
第一行package main是定义了我们此代码文件属于哪一个go package（我觉得可以近似理解为Python module，也就是作为一个独立的part进行开发的）
每个go的源代码文件，非注释第一行都必须给出一个定义，package xxx，是一个包的声明
package main 是一个特殊的包定义，每一个go应用程序（编译的时候）都一定要包含一个main包
main package代表着这个包是一个独立的、可执行的程序，可以近似理解为是一个主干
也可以套用汽车部件与汽车引擎的关系来近似理解
*/
package main

import (
	"fmt"

	"github.com/q1mi/hello"
) // 导入外部的包的时候，需要在go mod文件里面进行一个登记，其实mod文件就是管理了你这个项目的所有相关信息，可以用go mod tidy进行“刷新登记”

/*
imoprt 是一个导入包的操作，和R Python是一样的，不过要注意，go里面的引号一定是双引号
也可以通过github链接直接导入github的包 导入多个包用个括号把他们括起来就可以了
go语言也可以利用 go fmt xxx.go 来对go进行格式标准化，一些你写的不太规范的格式会给你自动更改，不过一般IDE在保存的时候会帮你自动做这个
*/

/*
fmt包是go的一个格式化输出的包，差不多就是一个基本包吧......
go fmt这个功能对于导入了没有用的包，会直接删掉的，go语言这点实在是牛逼
*/

/*
接下来就是定义函数并且写函数体内容，其实这里和R很像
main函数也是一个特殊的函数，和main package一样，main函数标志着该源代码开始执行
一般来说除了init函数以外（初始化），main函数是整个程序启动后运行的第一个函数
init可以应用于任意包中，且可以重复定义多个；main函数只能用于main包中，且只能定义一个
*/

// 像一些“没有实际行为”的语句，比如声明变量、函数、结构体，是可以放在函数体外面的
var n int

func main() { // go语言中花括号必须跟在东西后面，不能单独一行
	fmt.Println("hello it's my first go code") // fmt.Println这函数就是在终端进行输出，和print一样
	hello.SayHi()
	n = 100
	fmt.Println(n) // 上面对n赋值以后，如果没有像这样打印n，go会提醒你 n is used，感觉这个提示挺蠢的......（虽然没报错）

	// 涉及到代码逻辑（感觉可以理解为代码行为）的语句，必须放在函数内部，不能像Python那样随便放
}
// main包是比较特殊的, 编译之后是一个可执行文件, 如果一个文件没有main包和main函数, 是无法build成可执行函数的
package main

// 在加载包执行的时候, 一定要在go.mod文件下执行go run命令, 有点类似于node中在package.json目录下执行node命令
import (
	// 引入官方标准包, 存储在GOROOT目录中
	"fmt"

	// 引入第三方包, 存储在GOPATH目录中(会发现go-mod中引入了一个第三方依赖)
	"rsc.io/quote"

	// 引入本地本模块的其他包(这种写法不需要再go-mod中进行引入)
	"pogusanqian.com/project01/util"
	// TODO 引入本地其他模块的包
	// TODO 如果想要引入同一个包下的其他文件, 应该怎么引入?
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(util.Name)
}

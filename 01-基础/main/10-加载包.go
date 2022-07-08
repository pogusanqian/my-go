// main包是比较特殊的, 编译之后是一个可执行文件, 如果一个文件没有main包和main函数, 是无法build成可执行函数的
package main

// 在加载包执行的时候, 一定要在go.mod文件下执行go run命令, 有点类似于node中在package.json目录下执行node命令
import (
	// 引入官方标准包, 存储在GOROOT/src目录中
	"fmt"

	// 引入第三方包, 存储在GOPATH/pkg/mod目录中
	"rsc.io/quote"

	// 引入本地本模块的其他包(这种写法不需要再go-mod中进行引入, 另外这里的util即使包名又是路径名称)
	"pogusanqian.com/project01/util"

	// 引入本地其他模块的包(go mod edit -replace pogusanqian.com/greetings=../02-demo/greetings)
	"pogusanqian.com/greetings"
)

func main() {
	// fmt.Println(MyName)  // main包比较特殊, 这个包中就只有一个文件, 其他文件打不进来
	fmt.Println(quote.Go())
	fmt.Println(util.Name)
	fmt.Println(greetings.GetStudent())
}

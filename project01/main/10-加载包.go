package main

import (
	// 引入官方标准包, 存储在GOROOT目录中
	"fmt"
	// 引入第三方包, 存储在GOPATH目录中(会发现go-mod中引入了一个第三方依赖)
	"rsc.io/quote"
	// 引入本地本模块的其他包(这种写法不需要再go-mod中进行引入)
	"pogusanqian.com/project01/util"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(util.Name)
}

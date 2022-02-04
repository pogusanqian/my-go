package main

import (
	"fmt"
	// 这里的加载应该是同步加载, 先加载init文件, init文件加载完成之后才会执行本文件的init方法和main方法
	"pogusanqian.com/project01/util"
)

func init() {
	fmt.Println("main文件的init函数", util.Name)
}

func main() {
	fmt.Println("main文件的主函数", util.Name)
}

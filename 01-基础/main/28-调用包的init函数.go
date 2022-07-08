package main

import (
	"fmt"
	// 进行了同步加载
	"pogusanqian.com/project01/util"
)

func init() {
	fmt.Println("main文件的init函数", util.Name)
}

func main() {
	fmt.Println("main文件的主函数", util.Name)
}

package main

import (
	"fmt"
	"os"
)

// 获取命令行参数的时候, 执行的是build之后的执行文件 .\53.1-获取命令行参数.exe root 123123
func main() {
	fmt.Println("命令行的参数有", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
}

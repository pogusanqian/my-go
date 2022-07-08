package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 查看本地机器的逻辑CPU个数
	fmt.Println("老CPU数量", runtime.NumCPU())

	// GOMAXPROCS设置可同时执行的最大CPU数
	runtime.GOMAXPROCS(1)
	fmt.Println("新CPU数量", runtime.NumCPU())
}

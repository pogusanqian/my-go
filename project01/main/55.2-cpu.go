package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	// TODO 设置cpu(1.8之后, 默认设置的是多核), 感觉这个函数失效了, 不起作用
	runtime.GOMAXPROCS(4)
	fmt.Println(runtime.NumCPU())
}

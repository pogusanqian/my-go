package main

import (
	"fmt"
	"os"
)

// 流: 数据在文件和内存之间经历的路径; 输入流读文件, 输出流写文件(输入输出的主语是计算机)
func main() {
	file, err := os.Open("D://Code//my-go//01-基础//go.mod")
	if err != nil {
		fmt.Println("打开文件失败", err)
	}

	err = file.Close()
	if err != nil {
		fmt.Println("关闭文件失败", err)
	}
}

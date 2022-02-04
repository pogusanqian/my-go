package main

import (
	"fmt"
	"os"
)

/**
* 流: 数据在文件和内存之间经历的路径; 输入流读文件, 输出流写文件
 */
func main() {
	file, err := os.Open("D://Code//my-go//project01//go.mod")
	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	err = file.Close()
	if err != nil {
		fmt.Println("关闭文件失败", err)
	}
}

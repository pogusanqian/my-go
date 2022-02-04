package main

import (
	"fmt"
	"os"
)

/**
* os.Stat()返回值判断
* 	1. 如果返回的错误是nil, 说明文件或目录存在
*   2. 如果返回的错误使用os.IsNotExist()判断为true, 说明文件或目录存在
*   3. 如果是其他错误类型, 不确定是否存在
 */
func PathExits(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

func main() {
	flag, _ := PathExits("D:/abcs.txt")
	if flag {
		fmt.Println("文件存在")
	} else {
		fmt.Println("文件不存在")
	}
}

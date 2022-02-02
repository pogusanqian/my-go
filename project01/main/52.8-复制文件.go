package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	// 获取源文件
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	// 获取目标文件
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	// copy文件
	return io.Copy(writer, reader)
}

func main() {
	srcFile := "d:/flower.jpg"
	dstFile := "e:/abc.jpg"
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Printf("拷贝完成\n")
	} else {
		fmt.Printf("拷贝错误 err=%v\n", err)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打开一个存在的文件中，将原来的内容覆盖成新的内容10句 "你好，尚硅谷!"
	filePath := "d:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	str := "你好,尚硅谷!\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

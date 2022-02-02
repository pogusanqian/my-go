package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 以写和创建两种模式打开一个文件, 0666是用在linx下的, 设置文件权限
	file, err := os.OpenFile("D:/abc.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()

	str := "你好,世界\n" // 直接使用\n即可, window下的记事本现在都支持\n和\r\n两种换行
	// 写入时，使用带缓存的 *Writer, 最火需要使用flush将内存中的数据刷新到文件
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

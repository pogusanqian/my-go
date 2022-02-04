package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("D://Code//my-go//project01//go.mod")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer file.Close() //要及时关闭file句柄，否则会有内存泄漏.

	// 创建一个 *Reader  ，是带缓冲的; defaultBufSize = 4096 //默认的缓冲区为4096
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		if err == io.EOF {                  // io.EOF表示文件的末尾
			break
		}
		fmt.Printf(str)
	}
	fmt.Println("文件读取结束...")
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "d:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()

	str := "ABC,ENGLISH!\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

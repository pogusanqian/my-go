package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	// 建立和服务器的tcp连接
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	// 读取终端中的数据, os.Stdin代表标准输入[终端]
	reader := bufio.NewReader(os.Stdin)
	for {
		// 从终端读取一行用户输入, 并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}
		// 如果用户输入的是 exit就退出
		if strings.TrimSpace(line) == "exit" {
			fmt.Println("客户端退出..")
			break
		}

		//再将line 发送给服务器
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
	}
}

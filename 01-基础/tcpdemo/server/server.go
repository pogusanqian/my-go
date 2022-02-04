package main

import (
	"fmt"
	_ "io"
	"net" //做网络socket开发时,net包含有我们需要所有的方法和函数
)

func process(conn net.Conn) {
	defer conn.Close() //关闭conn
	for {
		buf := make([]byte, 1024)
		// 等待客户端通过conn发送信息; 如果客户端没有wrtie[发送]，那么协程就阻塞在这里(这里是不会发生死锁的); 返回的n表示客户通讯的数据字节
		n, err := conn.Read(buf)
		if err != nil { // 常见的错误有客户端主动断开连接
			fmt.Printf("err=%v", err)
			return
		}
		// 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听....")
	// 使用tcp协议监听本地的8888端口, 相当于启动了服务器
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("等待客户端来链接....")
		conn, err := listen.Accept() // 这里会阻塞等待, 只有和客户端建立连接之后才会继续往下走
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc con=%v 客户端IP=%v\n", conn, conn.RemoteAddr().String())
		}
		// 开启协程和客户端通信(这里需要把和客户端建立的coon连接传递过去)
		go process(conn)
	}
}

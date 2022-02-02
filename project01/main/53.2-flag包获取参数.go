package main

import (
	"flag"
	"fmt"
)

func main() {
	//定义几个变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	//&user 就是接收用户命令行中输入的 -u 后面的参数值, 默认值是空字符串
	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "p", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
	flag.IntVar(&port, "p", 3306, "端口号，默认为3306")

	// 转换方法(从命令行切片参数中转换), 必须调用
	flag.Parse()
	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}

package main

import (
	"fmt"
)

// 关闭管道后, 管道中的数据只能读取, 不能再往管道中写入了
func main() {

	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200

	// 关闭管道
	close(intChan)

	// 读数据可以成功, 写数据是失败的
	fmt.Println("n1=", <-intChan)
	intChan <- 300
}

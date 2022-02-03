package main

import (
	"fmt"
)

// 管道可以声明为双向管道(能读能写), 只读管道, 只写管道
func main() {

	// 双向管道, 能读能写
	chan1 := make(chan int, 3)
	fmt.Println(chan1)

	// 只写管道
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 20

	// 只读管道(数据从其他地方复制过来)
	var chan3 <-chan int
	// num2 := <-chan3 // 因为管道中没有数据, 会一直等待, 报deadlock错误
	fmt.Println("num2", chan3)
}

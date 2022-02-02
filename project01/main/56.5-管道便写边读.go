package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("writeData ", i)
	}
	close(intChan) // 关闭
}

// 管道遍历的时候时先将管道关闭, 是为了防止range迭代死锁; 但是使用for循环读取数据还是可以的, 注意再管道中写完数据时, 要将管道进行关闭就行了
func readData(intChan chan int, exitChan chan bool) {
	for {
		// 管道关闭后, 这里的ok只就是一个false, 表示管道已经关闭
		v, ok := <-intChan
		if !ok {
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	exitChan <- true
	close(exitChan)
}

func main() {

	//创建两个管道, 第二个管道时主线程继续执行的标志
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	// 使用协程进行边读边写
	go writeData(intChan)
	go readData(intChan, exitChan)

	// TODO 感觉这里使用for循环一致读取数据也不是很好, 不知道还有没有更好的解决方案
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}

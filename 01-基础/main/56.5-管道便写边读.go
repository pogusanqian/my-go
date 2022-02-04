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

/**
* 如果一个管道只有写, 没有读, 编译器发现管道进行了死锁; 例如我们要插入50条数据,但是管道只有10个长度, 后面四十个会发生阻塞, 进而演变成死锁
* 如果一个管道既有读有有写, 是不会发生死锁的; 如果写比读快, 当写到第11个时, 管道插入就会阻塞进行等待, 当读取走数据后才会进行插入; 这种阻塞是不会引起死锁的
 */
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

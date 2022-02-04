package main

import (
	"fmt"
)

// 向intChan管道中放入数据
func putNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {
		intChan <- i
	}
	close(intChan)
}

// 从intChan取出数据, 并判断是否为素数, 如果是, 就放入到primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		// 判断是否是素数
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag { // 放到素数管道
			primeChan <- num
		}
	}

	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
	// 往退出管道中插入数据
	exitChan <- true
}

func main() {
	// 定义三个管道
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000)
	exitChan := make(chan bool, 8)

	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	// 当我们从exitChan取出了4个结果, 就关闭prprimeChan管道; 如果从管道中取值时, 如果管道中没有值, 便会阻塞等待; 其实这里用len()判断管道的长度比较好
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	//遍历我们的 primeChan ,把结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("main线程退出")
}

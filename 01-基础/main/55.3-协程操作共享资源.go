package main

import (
	"fmt"
	"time"
)

var myMap = make(map[int]int, 10)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap[n] = res
}

// 报错: concurrent map writes: 并发操作公共资源map
func main() {
	// 开启200个线程计算阶乘
	for i := 1; i <= 20; i++ {
		go test(i)
	}
	// 由于主线程结束后, 协程也会强制结束, 所以这里将主线程进行了等待
	time.Sleep(time.Second * 5)
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}

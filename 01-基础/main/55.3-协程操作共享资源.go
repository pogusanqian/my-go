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

/**
* myMap属于全局变量, 20个协程同一时间操作这个变量, 就会报错: concurrent map writes(并发操作公共资源)
* go build -race ./main.go 打的exe文件能正常执行, 并输出有多少协程出现了竞争
*/
func main() {
	// 开启20个协程计算阶乘
	for i := 1; i <= 20; i++ {
		go test(i)
	}
	// 由于主线程结束后, 协程也会强制结束, 所以这里将主线程进行了等待
	time.Sleep(time.Second * 5)
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}

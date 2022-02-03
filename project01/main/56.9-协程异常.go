package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	// 这里没有make空间, 所以会报错
	var myMap map[int]string
	myMap[0] = "golang"
}

// 在Go中, 如果一个协程出现panic异常, 会导致其他异常以及主线程崩溃; 需要使用在协程中使用recover捕获异常
func main() {
	go sayHello()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}
}

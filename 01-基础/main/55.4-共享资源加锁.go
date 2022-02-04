package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁, lock是我们声明的变量, Mutex表示互斥锁
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	// 操作共享资源时, 进行了加锁和解锁
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

// 使用全局变量加锁是有一个问题的, 那就是主线程要等待多少秒
func main() {
	for i := 1; i <= 20; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 5)
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}

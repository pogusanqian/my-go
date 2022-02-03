package main

import (
	"fmt"
)

// 使用select可以解决从管道中读取数据, 因管道没有关闭, 而阻塞造成的死锁问题
func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	for {
		select {
		//注意:intChan一直没有关闭, 不会一直阻塞而deadlock,, 会自动到下一个case匹配
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%s\n", v)
		default:
			fmt.Printf("都取不到了，不玩了, 程序员可以加入逻辑\n")
			return // 这里不能使用break, 因为break不是跳出了for循环, 而是跳出了select; 可以采用label
		}
	}
}

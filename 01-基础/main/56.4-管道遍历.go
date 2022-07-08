package main

import (
	"fmt"
)

/**
* 1. 管道遍历的时候, 只支持使用range遍历, 如果使用普通的for循环遍历是不行的, 因为管道的长度会随着数据的取出递减
* 2. 遍历管道的时候需要先将管道close, 否则的话是会报错的; 因为遍历管道的时候, 如果管道没有关闭的话, 当将管道中的所有数据都遍历出来之后, range会仍然一致等待管道, 就会发生死锁
* 3. 在遍历关闭的管道时, 遍历完数据后, 会退出遍历
 */
func main() {
	intChan := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan <- i
	}
	
	close(intChan)
	for v := range intChan {
		fmt.Println("v=", v)
	}
}

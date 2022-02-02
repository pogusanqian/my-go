package main

import (
	"fmt"
)

/**
* 1. 管道本质是一个数据结构--队列, 数据先进先出(FIFO: firest in first ouit)
* 2. 管道可以实现数据安全, 不需要我们在手动加锁; 多个协程操作同一个管道时, 不会发生资源竞争
* 3. channel是有数据类型的, string类型的管道, 只能存放string类型
* 4. 管道是引用类型数据
* 5. 在没有使用协程的情况下, 如果将管道中的所有数据取出, 再去就会报错deadlock
* 6. 使用内置函数close关闭管道后, 管道就只能读数据不能再写数据了
 */
func main() {
	// 声明一个只能存放int的管道, 而且int的容量时3
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Println(intChan, &intChan)

	// 向管道中写入数据(注意, 这里最多只能存放3个数据, 超过容量之后是会报错的)
	intChan <- 10
	num := 10
	intChan <- num

	// 查看管道的长度和容量
	fmt.Println(len(intChan), cap(intChan))

	// 取出管道中的数据
	num1 := <-intChan
	num2 := <-intChan
	// num3 := <-intChan
	fmt.Println(num1, num2, len(intChan), cap(intChan))
}

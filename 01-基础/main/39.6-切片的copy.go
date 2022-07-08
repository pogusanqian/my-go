package main

import (
	"fmt"
)

/**
* copy拷贝值之后, 切片指向的数组还是同一个数组, 内存地址并没有变化
 */
func main() {
	arr := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	myslice1 := arr[:]
	myslice2 := []string{"张三", "李四", "王五"}
	fmt.Println(myslice1, &myslice1[0], &myslice1[4])

	// 将slice2的值copy到slice1中, 张三会把原本的值进行覆盖; 注意copy值之后, 修改slice2之后, 并不会影响slice1
	// 如果slice2有100个元素的话, 这里只会根据slice1的长度进行copy
	copy(myslice1, myslice2)
	myslice2[1] = "zhangsan"
	fmt.Println(myslice1, &myslice1[0], &myslice1[4])
	fmt.Println(myslice2)
}

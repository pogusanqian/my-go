package main

import (
	"fmt"
)

/**
* 内部排序: 将需要排序的数据加载到内存中进行排序(交换排序, 选择排序, 插入排序)
* 				 交换排序包含: 冒泡排序, 快速排序
* 外部排序: 数据量过大, 无法加载到内存, 需要借助外部存储进行排序(合并排序, 直接合并排序)
 */

func BubbleSort(arr *[5]int) {
	var temp = 0
	for i := 0; i < len(*arr)-1; i++ { // 定义循环的次数
		for j := 0; j < len(*arr)-1-i; j++ { // 里层for循环将最大的冒泡到最后
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
}

func main() {
	arr := [5]int{24, 69, 80, 57, 13}
	BubbleSort(&arr) // 传递过去的是一个指针数组
	fmt.Println(arr)
}

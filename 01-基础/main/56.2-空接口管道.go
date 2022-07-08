package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func main() {

	//定义一个存放任意数据类型的管道 3个数据
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom jack"
	allChan <- Cat{"小花猫", 4}

	//我们希望获得到管道中的第三个元素，则先将前2个推出
	<-allChan
	<-allChan
	newCat := <-allChan
	fmt.Println(len(allChan)) // 管道中的元素被取出后, 管道就成了空管道

	//从管道中获取的newCat据类型是空接口, 使用newCat.Name是不能通过编译的, 但是printf在运行的时候会打印出具体的类型
	fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat)
	// newCat.(Cat)进行了类型断言
	fmt.Printf("newCat.Name=%v", newCat.(Cat).Name)
}

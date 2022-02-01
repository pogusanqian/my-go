package main

import (
	"fmt"
	"pogusanqian.com/project01/util"
)

func main() {

	var midstu3 = util.MiddleStudent3{
		&util.Student3{ // 这里给Student3传递过去的是一个指针值
			Name: "张三",
			Age:  230,
		},
		222,
	}
	fmt.Println(midstu3)
	fmt.Println(*(midstu3.Student3), *midstu3.Student3) // 后者是对前者进行了简化, 去掉了()

}

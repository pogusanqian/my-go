package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//使用ioutil.ReadFile一次性将文件读取到位, 注意Open和close是封装到了ReadFile文件中的
	file := "D://Code//my-go//project01//go.mod"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	fmt.Printf("%s", content) // content是一个[]byte切片, 这里可以准换成string类型
}

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file1Path := "d:/abc.txt"
	file2Path := "d:/kkk.txt"
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("read file err=%v\n", err)
		return
	}
	// 这里使用的是一个file对象
	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("write file error=%v\n", err)
	}
}

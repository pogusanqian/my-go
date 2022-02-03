package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal kind=%v\n", rVal.Kind())
	//Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装
	rVal.Elem().SetInt(20)
}

func main() {
	var num int = 10
	reflect01(&num)          // 这里传递过去的是一个指针
	fmt.Println("num=", num) // 20
}

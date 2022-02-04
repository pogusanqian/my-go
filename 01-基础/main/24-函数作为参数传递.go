package main

import (
	"fmt"
)

type myInt int                           // 自定义myInt类型
type myFunType func(int, int) (int, int) // 自定义函数类型

func calc(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2
	return sum, sub
}

// func(int, int)(int, int)表示的是一个函数类型, 可以使用自定义类型替换
func calcWapper(myfun func(int, int) (int, int), num1 int, num2 int) int {
	res, _ := myfun(num1, num2)
	return res
}

func calcWapper2(myfun myFunType, num1 int, num2 int) int {
	res, _ := myfun(num1, num2)
	return res
}

func main() {
	res := calcWapper(calc, 20, 30)
	fmt.Println(res)

	res = calcWapper2(calc, 50, 30)
	fmt.Println(res)
}

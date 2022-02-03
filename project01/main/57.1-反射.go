package main

import (
	"fmt"
	"reflect"
)

/**
* 1. 反射可以在运行时动态获取变量的各种信息, 如变量的类型(Type), 列别(Kind); type和king当时基本类型的时候时相同的, 都是int; 当是结构体时, type是包名.Student, kind是struct
* 2. 如果变量时结构体变量, 还能获取到结构体本身的信息(字段, 方法)
* 3. 通过反射可以修改变量的值, 调用关联的方法
* 4. reflect包的TypeOf(), 可以获取变量的类型, 返回值是一个reflect.Type类型
* 5. reflect包的ValueOf(), 可以获取变量的值, 返回值是一个reflect.Value类型
* 6. 变量, interface{}, reflect.Value三者之间是可以相互转换的
 */
func reflectTest01(b interface{}) {
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2. 获取到 reflect.Value, 这个Value类型时不能进行运算的, 调用rval.Int()转换成int类型
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)

	// 4. 将rVal转换成空接口
	iV := rVal.Interface()

	// 将空接口转换成int类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

func main() {
	reflectTest01(100)
}

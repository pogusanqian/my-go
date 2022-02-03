package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {

	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	//3. 获取 变量对应的Kind, kind是一个常量值
	//(1) rVal.Kind() ==>
	kind1 := rVal.Kind()
	//(2) rTyp.Kind() ==>
	kind2 := rTyp.Kind()
	fmt.Printf("kind =%v kind=%v\n", kind1, kind2)

	//4. rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T \n", iV, iV)

	//5. 断言成Student类型
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}

}

type Student struct {
	Name string
	Age  int
}

func main() {
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest(stu)
}

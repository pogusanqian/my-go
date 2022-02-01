package main

import (
	"fmt"
)

/**
* 多态: 多态的实现的基础是接口; 形参如果是接口, 那个实参可以传递多种不同的类型
 */
type Usb interface {
	Start()
}

// 创建结构体, 并实现多态
type Phone struct {
}

type Camera struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}

// 电话方法是手机独有的
func (p Phone) Call() {
	fmt.Println("手机开始打电话。。。")
}

func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")
}

//计算机
type Computer struct {
}

func (c Computer) Working(usb Usb) { // 形参使用了接口, 可以实现多态
	usb.Start()
	// 类型断言调用call方法
	if phone, ok := usb.(Phone); ok == true {
		phone.Call()
	}
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)
}

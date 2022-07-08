package main

import (
	"fmt"
)

/**
* 1. Go中的接口是定义了一组方法, 并且没有实现; 接口可以实现多态和高内聚低耦合的思想
* 2. Go中的接口不能包含任何变量
* 3. Go中实现接口的时候, 不是使用implement接口
* 4. 接口不能创建实例
* 5. 自定义类型实现接口的时候, 需要将接口的所有方法都实现, 否则会报错; 我们的phone就没有是心啊usb2的test方法, 所以传参运行的时候, 就会报错
* 6. 自定义数据类型都能实现接口, 注意不单单时结构体, type myInt int自定义的数据类型也能实现接口
* 7. 接口之间是可以继承的
* 8. 接口类型是一个指针(引用类型), 如果没有进行初始化就使用的话, 会爆出nil
* 9. 空接口interface{}没有任何方法, 所以所有的类型都实现了空接口
*
* 继承价值: 解决代码的复用性和可维护性
* 接口价值: 在于设计规范, 让自定义的类型去实现这些方法
 */
type Usb interface {
	Start()
	Stop()
}

type Usb2 interface {
	Usb
	// Test()
}

// 创建结构体
type Phone struct {
}

type Camera struct {
}

// 结构体实现对应的接口; 注意在go中实现接口时, 只要指定方法就可以了, 现在Phone及实现了Usb接口和Usb2接口
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

// //计算机
// type Computer struct {
// }

// //编写一个方法Working 方法，接收一个Usb接口类型变量, 只要是实现了Usb接口的结构体, 都能进行传参
// func (c Computer) Working(usb Usb2) {
// 	usb.Start()
// 	usb.Stop()
// }

func main() {
	phone := Phone{}
	camera := Camera{}

	phone.Start()
	camera.Start()
}

package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

// 创建结构体
type Phone struct {
}

// 注意这里使用的是Phone指针实现的接口
func (p *Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p *Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

func main() {
	var p1 Phone = Phone{}
	// 这里依然会被编译器转换成(*p1).Start()
	p1.Start()
	p1.Stop()
}

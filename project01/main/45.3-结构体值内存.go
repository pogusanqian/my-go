package main

import (
	"fmt"
)

/**
* 结构体的属性值在内存中是连续存在的
 */
type Point struct {
	x int
	y int
}

type Rect struct {
	leftup, rightup Point
}

type Rect2 struct {
	left, right *Point
}

func main() {
	var stu = Point{25, 50}
	fmt.Println(&stu.x, &stu.y) // 0xc00000a168 0xc00000a170 ;int类型的就占用的是8个字节

	var rec1 = Rect{Point{20, 30}, Point{50, 100}}
	fmt.Println(&rec1.leftup.x, &rec1.leftup.y, &rec1.rightup.x, &rec1.rightup.x) // 0xc0000122c0 0xc0000122c8 0xc0000122d0 0xc0000122d0属性的属性值依然是连续的

	// 指针是连续的, 但是指针指向的地址不一定是连续的
	var rec2 = Rect2{&Point{20, 30}, &Point{50, 100}}
	fmt.Println(&rec2.left, &rec2.right)                                  // 0xc000058240 0xc000058248, 属性值是连续的
	fmt.Println(&rec2.left.x, &rec2.left.y, &rec2.right.x, &rec2.right.y) // 不一定是连续的
}

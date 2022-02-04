package main

import (
	"fmt"
)

func main() {
	var intArr = [5]int{0, 1, 2, 4, 5}
	myslice1 := intArr[:]  // 相当于[0:len(arr)]
	myslice2 := intArr[:3] // 相当于是[0:3]
	myslice3 := intArr[3:] // 相当于[3:len(arr)]

	fmt.Println(myslice1, myslice2, myslice3)
}

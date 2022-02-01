package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	// 等价与上面三局话
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {

	// Go自带的数组排序方法
	var intSlice = []int{0, -1, 10, 7, 90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄***"),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}

	//调用sort.Sort的自定义排序方法
	sort.Sort(heroes)
	for _, v := range heroes {
		fmt.Println(v)
	}
}

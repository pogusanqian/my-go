package main

import (
	"fmt"
	"log"
	"pogusanqian.com/greetings"
	"golang.org/x/example/stringutil"
)

func main() {
	names := []string{"张三", "李四", "王五"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	
	fmt.Println(stringutil.Reverse("Hello"))
	// fmt.Println(stringutil.ToUpper("Hello"))
}

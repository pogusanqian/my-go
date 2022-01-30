package main

import (
	"fmt"
	"time"
)

func main() {

	//1. 获取当前时间对象
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now)

	//2.通过now可以获取到年月日，时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month())) // 转换成了数字形式的
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//格式化日期时间
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dateStr=%v\n", dateStr)

	//格式化日期时间的第二种方式(2006-01-02 15:04:05是一个特殊标识, 相当于YYYY-MM-DD)
	fmt.Println(now.Format("2006-01-02 15:04:05")) // 获取当前的年月日时分秒
	fmt.Println(now.Format("2006---01---02"))      // 获取年月日
	fmt.Println(now.Format("15:04:05"))            // 获取时分秒
	fmt.Println(now.Format("2006"))                // 获取年丰

	// 常量, 0.1s只能使用time.Millisecond(毫秒) * 100表示, 不能使用tme.Second * 0.1表示
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond * 100)

	//Unix: 获取到秒的时间戳, UnixNano:获取到纳秒的时间戳
	fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n", now.Unix(), now.UnixNano())

}

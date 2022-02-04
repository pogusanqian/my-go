// 将sqlUtil文件打包到util包中, 这里不能随便写的, 因为sqlUtil.go文件就是在util目录下的
package util

import "fmt"

var Name = "张三"

func init() {
	fmt.Println("============Util的init函数")
}

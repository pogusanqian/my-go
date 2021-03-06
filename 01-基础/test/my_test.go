package test

import (
	"testing"
)

/**
* 1. 测试用例名必须以_test.go结尾, 因为这种结尾的文件, 会被go自带的test框架进行加载; 不需要写main方法的, go test会启动test框架, 这个框架会自动加载_test.go文件
* 2. 测试方法的形参必须是*testing.T, 这个对应封装了对应的测试方法
* 3. go test运行正确不会输出日志, go test -v 正确的日志和错误的日志都会进行输出
* 4. 测试的显示结果会显示函数执行的时间; Pass表示通过, Fail表示失败
* 5. go test -v my_test.go 运行指定的测试文件, go test -v 运行当前目录下的所有测试文件(不会遍历子目录)
* 6. go test -v -test.run TestAddUpper 只运行一个测试方法(-test.run是关键字)
* 7. go test运行是有缓存的, 输出的是上次运行的结果(实际上并没有运行代码, 因为我们的代码没有改变), 如果改变了代码, 运行的就不是缓存了
* 				go clean -testcache							清除缓存
*   			go test -v -count=1 ./**				忽略缓存执行
*
 */
func TestAddUpper(t *testing.T) {
	// 同一个包下的文件是可以直接调用的, 即使没有写在同一个文件中
	res := addUpper(10)
	if res != 55 {
		// 输出错误日志
		t.Fatalf("执行错误，期望值=%v 实际值=%v\n", 55, res)
	}
	// 输出正确日志
	t.Logf("成功执行函数")
}


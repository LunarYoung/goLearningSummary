package main

import "fmt"

func main() {
	//goto可以用在任何地方，但是不能跨函数使用

	fmt.Println("11111")
	goto End//直接跳到33333的输出执行语句

	fmt.Println("22222")// 此行被忽略
	End:
	fmt.Println("33333")
}

package main

import "fmt"

func main() {
	//test(1,4,5,6,5);
	test(1, 4, 8, 5, 3, 5, 2, 2, 5, 8, 4)
}

func test(args ... int64) {
	//MyFunc(args...)//args后面有...
	//从args[2]开始（包括本身），把后面所有参数传递过去
	MyFunc(args[2:]...)
	//从args[2]开始 （包括本身），到args[3]结束（不包括3），长度为3-2=1
	MyFunc(args[2:3]...)
	//从0开始，到2结束
	MyFunc(args[:2]...)
}

// ...int64类型这样的类型， ...type不定参数类型
func MyFunc(args ...int64) { //有其他类型参数，不定参数一定放到最后
	fmt.Println("args长度：", len(args))
	if (len(args) > 0) {
		fmt.Printf("其中:\n");
	}
	for i, arg := range args {
		fmt.Printf("第%d个参数是%d\n", i+1, arg)
	}
}

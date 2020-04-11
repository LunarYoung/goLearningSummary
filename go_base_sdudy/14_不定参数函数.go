package main

import "fmt"

func main() {
	MyFunc(1, 4, 5, 6, 5);
	MyFunc(1, 4, 8, 5, 3, 5, 2, 2, 5, 8, 4)
}

// ...int64类型这样的类型， ...type不定参数类型
func MyFunc(args ...int64) {
	fmt.Println("args长度：", len(args))
	if (len(args) > 0) {
		fmt.Printf("其中:\n");
	}
	for i, arg := range args {
		fmt.Printf("第%d个参数是%d\n", i+1, arg)
	}
}

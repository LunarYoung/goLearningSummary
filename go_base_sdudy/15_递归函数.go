package main

import "fmt"

type int int64

func main() {
	test(10);
}

func test(a int) {
	//函数调用自身
	if (a > 1) {
		test(a - 1)
	}
	fmt.Println("a=", a);
}

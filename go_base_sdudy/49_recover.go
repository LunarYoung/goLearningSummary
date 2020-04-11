package main

import "fmt"

func testb(x int) {
	defer func() { // 数组越界也不终止运行
		if err := recover(); err != nil { // 产生了panic异常
			fmt.Println(err);
		}
	}()

	var a [10]int;
	a[x] = x * x;
}

func main() {
	fmt.Println("aaa");
	testb(20);
	fmt.Println("bbb");
}

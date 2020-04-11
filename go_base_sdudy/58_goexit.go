package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("test:bbbb");
	//return; // 这里如果加return，则test:aaaa不会输出

	runtime.Goexit(); // 终止所在的子进程，后面的bbbb都不会输出
	fmt.Println("test:aaaa");
}

func main() {
	go func() {
		fmt.Println("aaaa");
		test();
		fmt.Println("bbbb");
	}();
	for {

	}
}

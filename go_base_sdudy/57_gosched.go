package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println("Hello World");
		}
	}();

	for j := 0; j < 2; j++ {
		// 让出时间片，先让别的协议执行，它执行完，再回来执行此进程
		runtime.Gosched();          // 这时候如果加个...，就...
		fmt.Println("Are You OK?"); // 主进程直接退出，子进程连运行的机会都没有
	}
}

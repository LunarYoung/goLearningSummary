package main

import (
	"fmt"
	"time"
)

func newFunc() {
	for {
		fmt.Println("子进程");
		time.Sleep(time.Second);
	}
}

func main() {
	// 子进程调用newFunc方法
	go newFunc();

	// 主进程执行
	i := 0;
	for {
		if (i == 10) {
			break; // 主进程退出后，主进程也会gg
		}
		fmt.Println("主进程");
		time.Sleep(time.Second);
		i++;
	}
}

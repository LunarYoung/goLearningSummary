package main

import (
	"fmt"
	"time"
)

func main() {
	//// 1.
	//timer := time.NewTimer(2 * time.Second); // 创建一个定时器，设置时间为2s，时间到了，只会响应一次
	//<-timer.C
	//// 2.
	//time.Sleep(2 * time.Second);
	//// 3.
	<-time.After(2 * time.Second); // 定时（阻塞）2s，2s后产生一个事件，往channel中写内容

	// 时间到了之后
	fmt.Printf("时间到");
}

func main0() {
	timer := time.NewTimer(1 * time.Second); // 创建一个定时器，设置时间为1s，时间到了，只会响应一次
	for {
		_, ok := <-timer.C
		if (ok == false) {
			break;
		}
		fmt.Printf("时间到\n");
	}
}

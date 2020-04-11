package main

import (
	"fmt"
	"time"
)

// 定时器重置
func main() {
	timer := time.NewTimer(3 * time.Second);
	go func() {
		<-timer.C
		fmt.Println("子协程可以打印了，因为定时器的时间到");
	}();

	// 定时器重置为1秒
	timer.Reset(1 * time.Second);
	for {

	}
}

// 定时器停止
func main0() {
	timer := time.NewTimer(3 * time.Second);
	go func() {
		<-timer.C
		fmt.Println("子协程可以打印了，因为定时器的时间到");
	}();

	// 计时器停止了
	timer.Stop();
	for {

	}
}

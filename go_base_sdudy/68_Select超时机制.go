package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int);
	quit := make(chan bool);

	// 新开一个协程
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println(num);
				break;
			case <-time.After(3 * time.Second): // 三秒钟读不出来数据，触发结束通道写入
				fmt.Println("超时");
				quit <- true;
				break;
			}
		}
	}();

	//
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second);
		ch <- i;
	}

	// 接收到quit信号，程序结束
	<-quit;
	fmt.Println("程序结束");
}

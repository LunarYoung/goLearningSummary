package main

import (
	"fmt"
	"time"
)

var (
	ch  = make(chan int)
	ch2 = make(chan string)
)

/**
 * goroutine 通过通信来共享内存，而不是通过共享内存来通信
 */

func Printer(str string) {
	for _, char := range str {
		fmt.Printf("%c", char);
		time.Sleep(time.Microsecond);
	}
}

func Printer1() {
	defer fmt.Println("打印子进程1执行完毕");

	//
	Printer("ABCDEFGHIJKLMNOPQRSTUVWXYZ");
	ch <- 666;
	// 给管道写数据，发送
}

func Printer2() {
	defer fmt.Println("打印子进程2执行完毕");

	//
	_ := <-ch; // 从管道取数据，如果通道没有数据他就会阻塞
	Printer("abcdefghijklmnopqrstuvwxyz");
	ch2 <- "工作完毕，请退出";
}

func main() {
	// 两个子进程同时调用一个函数
	go Printer1()
	go Printer2();

	// 没有数据前，阻塞
	x := <-ch2;
	fmt.Printf("%s", x);
}

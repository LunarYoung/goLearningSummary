package main

import (
	"fmt"
	"time"
)

func Printer(str string) {
	for _, char := range str {
		fmt.Printf("%c", char);
		time.Sleep(time.Microsecond);
	}
}

func main() {
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
	lower := "abcdefghijklmnopqrstuvwxyz";

	// 两个子进程同时调用一个函数
	go Printer(lower);
	go Printer(upper);
	for {

	}
}

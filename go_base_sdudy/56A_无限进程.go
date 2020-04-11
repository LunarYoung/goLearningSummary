package main

import (
	"fmt"
	"strconv"
	"time"
)

func newFunc(i int) {
	go newFunc(i + 1); // 调用子函数
	for {
		fmt.Println("子进程 - " + strconv.Itoa(i));
		time.Sleep(time.Second);
	}
}

func main() {
	newFunc(0);
}

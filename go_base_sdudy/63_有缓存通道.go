package main

import (
	"fmt"
	"time"
)

var (
	num   = 2
	count = 3
)

func main() {
	ch := make(chan int, num);

	//
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch));

	// 新建子进程
	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("子进程：i=%d\n", i);
			ch <- i; // 往chan写内容
			fmt.Printf("子进程：i=%d成功，len(ch)=%d，cap(ch)=%d\n", i, len(ch), cap(ch));
		}
		close(ch); // 关闭管道。关闭channel后无法再发送数据
	}();

	// 延时
	time.Sleep(time.Second * 2);

	//
	for i := 0; i < count; i++ {
		fmt.Printf("主进程：开始读数据\n");
		num, ok := <-ch
		if (ok == true) { // 如果ok为true，说明管道没有关闭
			fmt.Printf("主进程：num = %d\n", num);
		} else { // 反之，管道关闭
			break;
		}
	}

	//for j := range ch { //可以通过range遍历
	//	fmt.Printf("num = %d", j);
	//}
}

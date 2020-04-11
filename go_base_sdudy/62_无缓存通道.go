package main

import (
	"fmt"
	"time"
)

var (
	count = 3
)

func main() {
	// 创建一个无缓存的channel
	ch := make(chan int, 0);

	// len(ch)缓冲区剩余数据个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch));

	// 新建子进程
	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("子进程：i=%d\n", i);
			ch <- i; // 往chan写内容
			fmt.Printf("子进程：i=%d成功\n", i);
		}
	}();

	// 延时
	time.Sleep(time.Second * 2);

	//
	for i := 0; i < count; i++ {
		fmt.Printf("主进程：开始读数据\n");
		num := <-ch
		fmt.Printf("主进程：num = %d\n", num);
	}
}

//len(ch) = 0, cap(ch) = 0
//子进程：i=0
//---->子进程写入0成功
//主进程：开始读数据
//---->主进程读取0成功
//主进程：num = 0
//主进程：开始读数据
//子进程：i=0成功
//子进程：i=1
//---->子进程写入1成功
//子进程：i=1成功
//子进程：i=2【这里绝不可能写入2成功，因为1主进程还没取走】
//---->主进程读取1成功，子进程可能在此时写入2成功
//主进程：num = 1
//主进程：开始读数据
//---->主进程读取2成功
//主进程：num = 2
//子进程：i=2成功

package main

import "fmt"

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i * i * i;
	}
	close(ch);
}

func consumer(ch chan int) {
	for {
		i, ok := <-ch;
		if (ok == false) {
			break;
		}
		fmt.Println(i);
	}
}

func main() {
	//ch := make(chan int); // 创建一个双向的channel
	//
	//// 双向channel能隐式转换为单向，但单向无法转换为双向
	//var writeCh chan<- int = ch; // 只能写，不能读
	//var readCh <-chan int = ch; // 只能读，不能写
	//writeCh <- 666;
	////<-writeCh;

	//
	ch := make(chan int); // channel传参，引用传递

	// 生产者，生产数字，写入channel
	go producer(ch);

	// 消费者，从channel读取去数据，打印
	consumer(ch);
}

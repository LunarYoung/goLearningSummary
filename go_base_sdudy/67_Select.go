package main

import "fmt"

func main() {
	ch := make(chan int);
	quit := make(chan bool);
	// 消费者，从channel中读取内容
	// 新建协程
	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch;
			fmt.Println("", num);
		}
		// 可以停止
		quit <- true;
	}();

	// 生产者，产生数字，写入channel
	fibonacci1(ch, quit);
}

// 原生斐波那契数列
func fibonacci() {
	s := make([]int, 10);
	for i := 0; i < 10; i++ {
		if (i < 2) {
			s[i] = i;
		} else {
			s[i] = s[i-2] + s[i-1];
		}
	}

	fmt.Println(s);
}

func fibonacci1(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1;
	for {
		//监听channel数据的流动
		select {
		case ch <- x:
			//x, y = y, x+y;
			tmp := x + y;
			x = y;
			y = tmp;
			break;
		case flag := <-quit:
			fmt.Println("flag =", flag);
			return;
		}
	}
}

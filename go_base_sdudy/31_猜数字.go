package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	rawNum := createNum();
	var num int
	for i := 0; i < 10; i++ {
		fmt.Println("输入一个数字（范围1000~9999）：");
		fmt.Scan(&num);
		if num < rawNum {
			fmt.Println("小了，继续猜");
		} else if num > rawNum {
			fmt.Println("大了，继续猜");
		} else {
			fmt.Println("对了，这个数就是", rawNum);
		}
	}
	fmt.Println("很可惜，你没有猜到，这个数其实是", rawNum);
}

func createNum() int {
	rand.Seed(time.Now().UnixNano());
	var num int;
	num = rand.Intn(8999) + int(1000);
	return num;
}

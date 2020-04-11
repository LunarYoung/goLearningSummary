package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	sliceNum := createNum();
	sliceNumSlice := numSlice(sliceNum);
	for i := 0; i < 10; i++ {
		res := startGame(sliceNumSlice);
		if (res == true) {
			fmt.Println("恭喜猜对，使用回合：", i+1);
			return;
		}
	}
	fmt.Println("很遗憾，没有猜对，其实这个数是", sliceNum);
}

func startGame(sliceNumSlice []int) (res bool) {
	var keyNum int;
	for {
		fmt.Printf("输入一个四位数：");
		fmt.Scan(&keyNum);
		if keyNum > 999 && keyNum < 10000 {
			break;
		}
	}
	keyNumSlice := numSlice(keyNum);
	n := checkSlice(sliceNumSlice, keyNumSlice);
	if n == 4 {
		return true;
	} else {
		return false;
	}
}

func createNum() int {
	rand.Seed(time.Now().UnixNano());
	var num int;
	num = rand.Intn(8999) + int(1000);
	return num;
}

func numSlice(num int) ([]int) {
	slice := make([]int, 4);
	slice[0] = num / 1000;
	slice[1] = num % 1000 / 1000;
	slice[2] = num % 100 / 10;
	slice[3] = num % 10;
	return slice;
}

func checkSlice(numSlice []int, keySlice []int) (n int) {
	n = 0;
	for i := 0; i < 4; i++ {
		if keySlice[i] < numSlice[i] {
			fmt.Printf("第%d个数小了\n", i);
		} else if keySlice[i] > numSlice[i] {
			fmt.Printf("第%d个数大了\n", i);
		} else {
			fmt.Printf("第%d个数猜对了\n", i);
			n++;
		}
	}
	fmt.Printf("\n");
	return n;
}

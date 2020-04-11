package main

import "fmt"

//目标：实现1+2+3+...+100

func main() {
	sum := add(100) //迭代方法
	fmt.Println(sum);
	sum = add2(100) //递归方法
	fmt.Println(sum)
}

func add(max int) (sum int) {
	for i := 1; i <= max; i++ {
		sum += i;
	}
	return;
}

func add2(max int) (sum int) {
	if (max == 1) {
		return 1;
	}
	return max + add2(max-1);
}

package main

import "fmt"

//_包名 不直接使用包里的函数，而是调用包里面的init函数

func init() { // 先于main执行
	fmt.Println("前述")
}

func main() {
	fmt.Println("正文")
}

// 函数执行顺序：包->init->main

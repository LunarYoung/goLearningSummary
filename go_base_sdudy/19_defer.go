package main

import "fmt"

//defer 延迟调用，只可用于函数中，函数结束前调用
func main() {
	fmt.Println("a")
	fmt.Println("bb")

	defer test(2) //后被执行

	fmt.Println("ccc")

	defer test(1) //先被执行

	fmt.Println("dddd")
	// defer函数发生错误，函数依旧会被执行
}

func test(c int) {
	for i := 0; i < c; i++ {
		fmt.Printf("%d", c)
	}
	fmt.Printf("\n")
}

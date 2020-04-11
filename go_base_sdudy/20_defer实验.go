package main

import "fmt"

func main() {

	a, b := 10, 20 //1 初始化

	defer func() { //4 被调用
		fmt.Printf("a=%d,b=%d\n", a, b)
	}()

	defer func(a int, b int) { //3 被调用(但值在被赋值前传入，所以还是旧值)
		fmt.Printf("a=%d,b=%d\n", a, b)
	}(a, b)

	a = 100
	b = 200 //2 重新被赋值
}

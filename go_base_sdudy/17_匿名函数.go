package main

import "fmt"

func main() {
	add := func(a int, b int) int {
		return a + b;
	}

	a, b := 10, 20
	c := add(a, b)

	fmt.Println("c=", c)

	func() {
		fmt.Printf("a=%d,b=%d\n", a, b)
	}() //有括号相当于调用

	func(a int, b int) {
		fmt.Println("c=", a+b)
	}(a, b) //定义的同时传参，直接调用

	max := func(a int, b int) int {
		//max:=a>b?a:b; 三目运算符不被允许使用
		var max int;
		if (a > b) {
			max = a;
		} else {
			max = b;
		}
		return max
	}
	fmt.Println("max number is:", max(a, b))
}

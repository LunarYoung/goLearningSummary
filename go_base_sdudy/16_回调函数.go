package main

import "fmt"

//定义个形式，形参要用
type FuncOne func(int, int) int

//具体形式，真正要计算的函数
func Add(a int, b int) int {
	return a + b;
}

func Minus(a int, b int) int {
	return a - b;
} //mult.div

func Mult(a int, b int) int {
	return a * b;
}

func Div(a int, b int) int {
	//float64(a)
	//return float64(a)/float64(b);
	return a / b;
}

func main() {
	a := 100
	b := 10
	fmt.Printf("a+b=%d\n", Calc(a, b, Add))
	fmt.Printf("a-b=%d\n", Calc(a, b, Minus))
	fmt.Printf("a*b=%d\n", Calc(a, b, Mult))
	fmt.Printf("a/b=%d\n", Calc(a, b, Div))
}

//回调函数，函数有一个参数是函数类型，这个函数就是回调函数
//计算器，可以进行四则运算
//多态，多种形态，调用同一个接口，不同的表现，可以实现不同表现，加减乘除
func Calc(a int, b int, calc FuncOne) (res int) {
	//fmt.Println("calc result:");
	res = calc(a, b);
	return res
}

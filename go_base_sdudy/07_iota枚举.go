package main

import "fmt"

func main() {
	const (
		a = iota
		b = iota
		c = iota
	)

	const d=iota//重置为零

	//常量自动生成，每隔一行自动加一
	//仅作常量赋值使用
fmt.Println(a,b,c)//output 0 1 2
//常量可以声明而不使用
}
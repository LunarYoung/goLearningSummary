//1) go语言以包作为管理单位
//2) 每个文件必须先声明包
//3) 程序必须有一个main包

package main

import "fmt"

func main()  {//左括号必须和函数名同行
	fmt.Println("Hello World")//Println会自动换行
	fmt.Printf("Hello World")//go语言结尾可以没有分号
}
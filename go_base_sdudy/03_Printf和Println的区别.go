package main

import "fmt"

func main() {
	str := "Hello, Year"
	year := 2017
	fmt.Printf("%s%d\n", str, year)
	fmt.Println(str, year) //变量之间自带空格
	//输出结果
	//Hello, Year2017
	//Hello, Year 2017
}

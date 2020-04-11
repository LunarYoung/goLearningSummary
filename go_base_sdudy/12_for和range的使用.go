package main

import "fmt"

func main() {
	str:="Today is Sunday"
	for i:=0; i<len(str);i++  {//len用于得出字符串长度和数组的元素个数
		fmt.Printf("第%2d个字符是：%c\n", i,str[i])
	}
	for j,c:=range str{
		fmt.Printf("第%2d个字符是：%c\n", j, c)
	}
}

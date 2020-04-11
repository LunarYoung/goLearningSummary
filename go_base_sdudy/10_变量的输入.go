package main

import "fmt"

func main()  {
	var a int
	fmt.Printf("请输入变量：")

	//fmt.Scanf("%d", &a)别忘了有个地址符号（和C语言一样）
	fmt.Scan(&a)//自动地

	fmt.Println("您输入的变量值为：",a)
}

package main

import "fmt"

func main()  {
	//交换变量
	var a,b int=10,20
	fmt.Printf("%d %d\n",a,b)//output 10 20
	a,b=b,a
	fmt.Printf("%d %d\n",a,b)//output 20 10

	tmp,_:=a,b//_是匿名变量，被设计就是要丢弃的
	fmt.Printf("%d", tmp)
}
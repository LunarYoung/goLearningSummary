package main

import "fmt"

func main() {
	var a int32=110
	fmt.Printf("%d\n", a)
	fmt.Printf("%v\n", &a)

	var p *int32=&a;
	fmt.Printf("%v\n", p)//p和a所在的地址一样

	(*p)++// 对p所指的值操作等价于对a操作
	fmt.Printf("%d\n", a)

	var p1,p2 *int32
	fmt.Println(p1,p2)//默认值 nil
	//(*p1)=2//err panic,因为p1没有合法指向

	p1=new(int32)
	p2=new(int32)
	fmt.Println(p1==p2)

	(*p1)=2//p1有了合法指向，赋值成功
	fmt.Println(*p1)
}

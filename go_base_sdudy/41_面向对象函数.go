package main

import (
	"fmt"
)

type People struct {
	name string
	sex  byte
	age  int
}

// 带有接收者的函数叫方法，接收者本身不能是指针（这里的接收者是People对象）
func (p People) output() {
	fmt.Println("姓名:", p.name);
	fmt.Println("性别:", p.sex);
	fmt.Println("年龄:", p.age);
}

//// 接收者类型不一样，函数名相同也不冲突
//func (tmp char)test()  {
//
//}
//func (tmp byte)test()  {
//
//}

func main() {
	// 定义同时初始化
	p := People{"Bob", 'm', 19};
	p.output();
	var p1 *People = &p; // 但是这里可以
	p1.output();
}

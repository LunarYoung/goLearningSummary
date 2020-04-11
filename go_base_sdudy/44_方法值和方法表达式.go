package main

import "fmt"

type People struct {
	name string
	sex  byte
	age  int
}

func (p People) output() {
	fmt.Println("姓名:", p.name);
	fmt.Println("性别:", p.sex);
	fmt.Println("年龄:", p.age);
}

// 接收者为People类对象的指针
func (p *People) addAgeP(a int) {
	p.age += a;
}

func main() {
	p := People{"Bob", 'm', 18};
	output := p.output; // 这个就是方法值，调用函数时，无需再传递接收者，隐藏了接收者
	output();
	f := (*People).addAgeP; // 方法表达式
	f(&p, 2);               // 只需要把接收者传递过去
	output();// 这里为什么不是20？？？
}

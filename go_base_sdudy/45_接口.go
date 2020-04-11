package main

import "fmt"

type Humaner interface {
	sayHi();
	sayYes();
}

type People struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	People
	id   int
	addr string
}

func (p *People) sayHi() {
	fmt.Printf("Hi, %s\n", p.name);
}

func (p *People) sayYes() {
	fmt.Printf("yes, %s\n", p.name);
}

func (p *Student) sayHi() {
	fmt.Printf("Hi, %d-%s\n", p.id, p.name);
}

func (p *Student) sayYes() {
	fmt.Printf("yes, %d-%s\n", p.id, p.name);
}

// 多态化调用，变相多态
func whoSayHi(i Humaner) {
	i.sayHi()
}

func main() {
	// 定义接口类型的变量
	var i Humaner;

	// 只要实现此接口方法的类型，那么这个变量的类型（接收者类型）就可以给i赋值
	p := &People{"Bob", 'm', 19};
	i = p;
	i.sayHi(); // 依照接收者类型选择不同的方法
	i.sayYes();
	s := &Student{People{"Ann", 'f', 17}, 1, "河北省保定市"};
	i = s;
	i.sayHi();
	i.sayYes();
	whoSayHi(p);
	whoSayHi(s);
	x := make([]Humaner, 2);
	x[0] = p;
	x[1] = s;
	for _, i := range x {
		i.sayHi();
	}
}

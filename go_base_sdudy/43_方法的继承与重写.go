package main

import "fmt"

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

func (p People) output() {
	fmt.Println("姓名:", p.name);
	fmt.Println("性别:", p.sex);
	fmt.Println("年龄:", p.age);
}

func (s Student) output() { // 这将重写People的output方法
	// 就近原则，先找本作用域的方法，找不到再找继承的方法
	fmt.Println("姓名:", s.name);
	fmt.Println("性别:", s.sex);
	fmt.Println("年龄:", s.age);
	fmt.Println("学号:", s.id);
	fmt.Println("地址:", s.addr);
}

func main() {
	s := Student{People{"Bob", 'm', 18}, 1, "Hebei Province,China"}; // Student可以使用People的方法
	s.output();
}

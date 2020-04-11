package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	// 每个成员都必须按顺序初始化
	var s1 Student = Student{1, "Angle", 'f', 18, "河北省石家庄市"};

	// 可指定成员初始化
	var s2 Student = Student{
		id:   2,
		name: "Bob",
		age:  19};
	fmt.Println("s1:", s1);
	fmt.Println("s2:", s2);

	echo(s1);
}

func echo(s Student)  {
	fmt.Println("id:", s.id);
	fmt.Println("name:", s.name);
	fmt.Println("sex:", s.sex);
	fmt.Println("age:", s.age);
	fmt.Println("addr:", s.addr);
}
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

// 接收者为People类对象
func (p People)addAgeV()  {
	p.age++;
}

// 接收者为People类对象的指针
func (p *People)addAgeP()  {
	p.age++;
}

func main() {
	p := People{"Bob", 'm', 19};
	p.output();
	p.addAgeV();
	p.output();
	p.addAgeP();
	p.output();
}

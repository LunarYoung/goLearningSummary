package main

import "fmt"

type PeopleInterface interface {
	sayHi();
}

type StudentInterface interface {
	PeopleInterface
	sing(lrc string);
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

func (p *Student) sayHi() {
	fmt.Printf("Hi, %d-%s\n", p.id, p.name);
}

func (p *Student) sing(lrc string) {
	fmt.Printf("学生%d在唱歌：%s\n", p.id, lrc);
}

func main() {
	var pi PeopleInterface;
	var si StudentInterface;
	s := &Student{People{"Mike", 'm', 20}, 1, "河北省保定市"};
	si = s;
	si.sing("啦啦啦");

	// 超集可以转换为子集
	pi = si;
	pi.sayHi();// People没有sayHi，但是超集（子类）有
	// 子集不能转换为超集
	//si = pi;
}

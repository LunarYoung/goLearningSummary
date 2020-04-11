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
	s1 := Student{1, "mike", 'm', 18, "bj"};
	fmt.Println(s1);
	reset(s1);
	fmt.Println(s1);
	reset1(&s1);
	fmt.Println(s1);
}

func reset(s Student) { // 形参无法改变实参，因为这里是值传递（数组可以，因为数组是址传递）
	s.id = 0;
	s.name = "-";
	s.sex = '-';
	s.age = 0;
	s.addr = "-";
}

func reset1(s *Student) { // 这里是址传递
	s.id = 0;
	s.name = "-";
	s.sex = '-';
	s.age = 0;
	s.addr = "-";
}

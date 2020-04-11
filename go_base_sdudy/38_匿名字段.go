package main

import "fmt"

type People struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	People // 只有类型，没有名字
	id   int
	addr string
}

func main() {
	// 顺序初始化，自动推导类型
	n1 := Student{People{"Bob", 'm', 22}, 1, "河北省石家庄市长安区"}

	// +v显示更详细的信息
	fmt.Printf("%+v\n", n1)

	// 指定成员初始化，没有初始化的成员自动赋值为0
	s1 := Student{People: People{age: 22}, id: 1} // 这时候，得加People:否则编译不过
	fmt.Printf("%+v\n", s1)

	s1 = n1

	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr);
	fmt.Println(s1.People.name, s1.People.sex, s1.People.age, s1.id, s1.addr);
}

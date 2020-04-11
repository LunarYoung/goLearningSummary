package main

import "fmt"

type People struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	*People // 只有类型，没有名字
	id   int
	addr string
}

func main() {
	s1 := Student{&People{"Bob", 'm', 22}, 1, "河北省石家庄市长安区"}
	fmt.Printf("%+v\n", s1)
	fmt.Printf("%+v\n", s1.People)
	fmt.Println(s1.name, s1.sex, s1.age)

	s1.People=&People{"Mike", 'm', 21}
	fmt.Println(s1.name, s1.sex, s1.age)
}

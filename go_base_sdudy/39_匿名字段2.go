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
	name string// 和People里的name的属性相重
}

func main() {
	n1 := Student{People{"Bob", 'm', 22}, 1, "河北省石家庄市长安区", "Bob1"}

	fmt.Printf("%+v\n", n1)
	fmt.Println(n1.People.name, n1.name)
}

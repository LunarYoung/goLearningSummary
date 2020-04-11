package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person
	id   int
	addr string
}

func main() {
	s1:=Student{Person{"mike", 'm', 24}, 1, "河北省石家庄市"}
	fmt.Printf("%+v", s1)
}

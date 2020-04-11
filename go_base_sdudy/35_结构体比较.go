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
	s2 := Student{1, "mike", 'm', 18, "bj"};
	s3 := Student{1, "mike", 'm', 19, "bj"};
	fmt.Println(s1 == s2);
	fmt.Println(s1 == s3);
	s3 = s2; // 同类型的两个结构体可以相互赋值
	fmt.Println(s1 == s3);
}

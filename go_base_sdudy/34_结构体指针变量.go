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
	var s1 *Student = &Student{1, "Angle", 'f', 18, ""};
	fmt.Println("s1:", *s1);
}

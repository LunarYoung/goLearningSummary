package main

import (
	"errors"
	"fmt"
)

func testa() {
	fmt.Println("aaaaaaaaaaaaaaa");
}

func testb() {
	fmt.Println("bbbbbbbbbbbbbbb");
}

func zuosi() { // 这里必然错误，因为数组越界了
	var a [10]int;
	a[20] = 11;
}

func main() {
	err1 := fmt.Errorf("%s", "this is a normal error1");
	fmt.Println("err1=", err1);
	err2 := errors.New("this is a normal error2");
	fmt.Println("err2=", err2);

	// 显示调用
	//testa();
	//panic("this is a panic test"); // 这里直接终止运行
	//testb();

	zuosi();
}

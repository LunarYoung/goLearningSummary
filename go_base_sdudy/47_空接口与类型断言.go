package main

import "fmt"

func main() {
	var i interface{} = 1;
	fmt.Println("i = ", i);
	// 空接口相当于任意一个类型

	s := make([]interface{}, 3);
	s[0] = 1;
	s[1] = "1";
	s[2] = byte('2');

	// if实现类型断言
	//for i, data := range s {
	//	fmt.Println(i, data);
	//	// 第一个返回结果的值，第二个返回结果的真假
	//	if value, ok := data.(int); ok == true {
	//		fmt.Printf("x[%d]为int，内容为%d\n", i, value);
	//	} else if value, ok := data.(string); ok == true {
	//		fmt.Printf("x[%d]为string，内容为%s\n", i, value);
	//	} else if value, ok := data.(byte); ok == true {
	//		fmt.Printf("x[%d]为byte，内容为%d\n", i, value);
	//	}
	//}

	// switch实现类型断言
	for i, data := range s {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d]为int，内容为%d\n", i, value);
		case string:
			fmt.Printf("x[%d]为string，内容为%s\n", i, value);
		case byte:
			fmt.Printf("x[%d]为byte，内容为%d\n", i, value);
		}
	}
}

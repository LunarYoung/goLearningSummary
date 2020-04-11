package main

import "fmt"

func main() {
	var id [50]int;
	for i := 0; i < len(id); i++ {
		id[i] = i + 1;
	}
	for key, value := range id {
		fmt.Printf("id[%d] = %d\n", key, value);
	}

	//var a [5]int = [5]int{1, 3, 4, 5, 9};// 全部初始化
	var a [5]int = [5]int{1, 3, 4}; // 部分初始化，没有初始化的元素，自动赋值为0
	for key2, value2 := range a {
		fmt.Printf("a[%d] = %d\n", key2, value2);
	}
}

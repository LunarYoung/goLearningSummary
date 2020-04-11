package main

import "fmt"

func main() {
	oldSlice := make([]int, 0, 1);
	newSlice := make([]int, 0, 1);
	for i := 0; i < 20; i++ {
		newSlice = append(oldSlice, i); // 二倍扩容
		fmt.Printf("%d:%d ===> %d:%d\n", len(oldSlice), cap(oldSlice), len(newSlice), cap(newSlice));
		oldSlice = newSlice;
	}

	a := []int{0, 1, 2, 3};
	b := []int{5, 6, 7, 8, 9, 10};
	fmt.Println("a = ", a);
	fmt.Println("b = ", b);
	copy(b, a); // 将b中的数组元素替换为a的相同位置
	fmt.Println("a = ", a);
	fmt.Println("b = ", b);
}

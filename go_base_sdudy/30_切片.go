package main

import "fmt"

func main() {
	array := [5]int{1, 3, 5, 7, 9};
	slice := array[1:3:5]; // 左闭右开(low,high,max)
	fmt.Println("slice = ", slice);
	// 不能直接赋值
	slice = append(slice, 2);
	fmt.Println("slice = ", slice);
	slice = append(slice, 3);
	slice = append(slice, 3);
	slice = append(slice, 3); // 长度被扩展
	fmt.Println("slice = ", slice);
	/**
	长度 high-low
	容量 max-low
	 */
	fmt.Printf("len=%d,cap=%d\n", len(slice), cap(slice));
	slice2 := make([]int, 3)
	fmt.Println("slice = ", slice2);
	fmt.Printf("len=%d,cap=%d\n", len(slice2), cap(slice2));
	s1 := array[:3] // 从0开始取3个元素，容量和原数组相同
	fmt.Printf("len=%d,cap=%d\n", len(s1), cap(s1));
	s2 := array[3:] // 从3开始取到末尾，容量和长度相同
	fmt.Printf("len=%d,cap=%d\n", len(s2), cap(s2));
}

package main

import "fmt"

func main() {
	var a [3][4]int;
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			a[i][j] = (i + 1) * (j + 1);
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j]);
		}
	}

	//var b [3][2]int=[3][2]int{
	//	{1,2},
	//	{2},
	//	{3,4}};// 最右边的大括号不能独立成一行
	var b [3][2]int = [3][2]int{1: {1, 1}}
	for _, value := range b {
		for _, val := range value {
			fmt.Printf("%d ", val);
		}
	}
}

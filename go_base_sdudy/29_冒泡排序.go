package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	var array [20]int = mkRandArray();
	fmt.Println("排序前 array = ", array);
	sort(&array);
	fmt.Println("排序后 array = ", array);
}

func mkRandArray() [20]int {
	var array [20]int;
	for i := 0; i < len(array); i++ {
		array[i] = mtRand(0, 100);
	}
	return array;
}

func mtRand(min int, max int) int {
	var len int = max - min;
	rand.Seed(time.Now().UnixNano());
	a := rand.Intn(len) + min;
	return a;
}

func sort(array *[20]int) {
	// 饱和冒泡排序(复杂度为 n-1^2)
	//for i := 0; i < len(array)-1; i++ {
	//	for j := 0; j < len(array)-1; j++ {
	//		if (array[j] > array[j+1]) {
	//			array[j], array[j+1] = array[j+1], array[j];
	//		}
	//	}
	//}
	// 不饱和冒泡排序(复杂度为 (n-1)+...+1  n(n-1)/2)
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if (array[j] > array[j+1]) {
				array[j], array[j+1] = array[j+1], array[j];
			}
		}
	}
}

package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2); // 单核和多核，结果可是不一样的呢
	for i := 0; i < 100000; i++ {
		if (i%200 == 0 && i > 0) {
			fmt.Printf("\n");
		}
		go fmt.Print(0);
		fmt.Print(1);
	}
}

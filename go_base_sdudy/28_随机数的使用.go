package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	//rand.Seed(666)
	rand.Seed(time.Now().UnixNano()); // 以时间做种子，可以使种子随机

	for i := 0; i < 5; i++ {
		fmt.Println("rand = ", rand.Int())     // 种子一样，则随机数一样
		fmt.Println("rand = ", rand.Intn(100)) // 限制在100以内，即0到99
	}
}

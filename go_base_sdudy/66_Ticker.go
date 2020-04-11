package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(3 * time.Second);
	for i := 0; ; i++ {
		<-ticker.C
		fmt.Println("i =", i);
		if (i == 5) {
			ticker.Stop();
			break;
		}
	}
}

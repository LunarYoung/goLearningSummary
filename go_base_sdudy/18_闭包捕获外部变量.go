package main

import "fmt"

func main() {
	a := 10
	str := "i love you"

	func() { //闭包以引用方式进行传递
		a = 60
		str = "do you love me?"
	}()

	fmt.Printf("a=%d,str=%s", a, str)
}

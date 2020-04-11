package main

import "fmt"

func main()  {
	num:=2
	switch num {
	case 1:fmt.Printf("这是%d", num)// 和C不同，这里自带break如果想要继续执行，请用
	case 2:fmt.Printf("这是%d", num)
	//fallthrough
	}
}

package main

import "fmt"

func main() {
	////支持初始化语句和变量本身
	//switch num:=1;num {
	//case 1:fmt.Printf("这是1")
	//case 2:fmt.Printf("这是2")
	//}
	score:=85
	switch {
	case score>90:fmt.Printf("你很优秀")
	case score>75:fmt.Printf("你很不错")
	case score>=60:fmt.Printf("你还可以")
	default:fmt.Printf("什么玩意儿")
	}
}

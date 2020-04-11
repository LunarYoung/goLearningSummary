//go可以返回多个值

package main
import "fmt"


func main()  {
	a,b,c:=getVar()
	fmt.Println(a,b,c)
}


func getVar() (int,int,int)  {//这里返回三个值
	return 1,2,3
}
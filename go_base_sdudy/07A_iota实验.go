package main

import "fmt"

func main()  {

	const a,b,c  = iota,iota,iota//iota在同一行中值都一样

	fmt.Println(a,b,c)
}

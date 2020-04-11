package main

import "fmt"

func main()  {
	//var a int64;
	//var b float64;

	var (
		a int64
		b float64
	);

	fmt.Printf("%d %f\n", a,b);

	//const c =10
	//const d float64=3.14

	const (
		c=10
		d float64=3.14
	)


	//可以自动推导类型
	//const (
	//	c=10
	//	b=3.14
	//)

	fmt.Println(c,d);
}
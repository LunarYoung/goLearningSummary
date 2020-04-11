package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var res bool;

	// Contains 字符串A是否包含字符串B（相当于php中的strstr）
	str := "My name is Bob";
	res = strings.Contains(str, "name");
	fmt.Println(res);

	// Join 字符串组合（相当于php中的join和implode）
	strs := []string{"My", "name", "is", "Bob"};
	str = strings.Join(strs, "#");
	fmt.Println(str);

	// Split 以指定分隔符拆分字符串（相当于php中的explode）
	strs2 := strings.Split(str, "#");
	fmt.Println(strs2);

	// Index（相当于php中的strpos）
	i := strings.Index(str, "name"); // 不包含子串返回-1，（php中不包含返回false）
	fmt.Println(i);

	// Repeat
	str2 := "Bob ";
	str2 = strings.Repeat(str2, 4);
	fmt.Println(str2);

	// Trim 去掉两头字符（相当于php中的trim，但没有默认的五空符）
	str2 = strings.Trim(str2, " ");
	fmt.Println(str2); // 字符串中最后的空格没有了

	// Fields 去掉空格之后，把元素放入切片中
	str3 := "        Are you ok?         ";
	strs3 := strings.Fields(str3);
	fmt.Println(strs3);

	//

	// Append
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true);    // 追加一个bool类型的值
	slice = strconv.AppendInt(slice, 1234, 10); // 追加一个十进制int的数值
	slice = strconv.AppendQuote(slice, "abcdefg");
	fmt.Println(slice);

	// Format
	var str4 string;
	str4 = strconv.FormatBool(false);
	fmt.Println(str4);
	str4 = strconv.FormatFloat(3.14159265, 'f', -1, 64); // 'f'指打印格式，-1指小数点位数（紧缩模式），64指float64处理
	fmt.Println(str4);
	str4 = strconv.Itoa(3402);
	fmt.Println(str4);

	// Parse
	var flag bool;
	var err error;
	flag, err = strconv.ParseBool("true");
	if err == nil {
		fmt.Println("flag =", flag);
	} else {
		fmt.Println("err =", err);
	}

	a, _ := strconv.Atoi("3566");
	fmt.Println(a + 1);
}

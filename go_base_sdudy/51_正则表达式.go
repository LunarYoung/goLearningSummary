package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf1 := "abc azc a7z 888 a9c tac";
	reg1 := regexp.MustCompile(`a.c`); // . 匹配除\n外的任意字符， 反引号`代表原生字符串
	if reg1 == nil {
		fmt.Println("error=", reg1);
		return;
	}

	result1 := reg1.FindAllStringSubmatch(buf1, -1); // n代表匹配数量，小于0代表无限匹配
	fmt.Println("result1=", result1);

	//
	buf2 := "43.14 567 adsgs 1.23 7. 8.99 1fasdfsa 6.66";
	//reg2 := regexp.MustCompile(`\d\.\d`);// 一位数字+"."+一位数字
	//reg2 := regexp.MustCompile(`\d\.\d+`);// 一位数字+"."+至少一位数字
	reg2 := regexp.MustCompile(`\d+\.\d+`); // 至少一位数字+"."+至少一位数字
	if reg2 == nil {
		fmt.Println("error=", reg2);
	}

	result2 := reg2.FindAllStringSubmatch(buf2, -1);
	fmt.Println("result2=", result2);
}

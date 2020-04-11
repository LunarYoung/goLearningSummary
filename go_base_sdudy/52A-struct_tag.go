package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	// 结构体变量一定要大写，否则获取不到
	Company  string   `json:"-"`        // 此字段不会输出到屏幕
	Subjects []string `json:"subjects"` // 二次编码
	IsOK     bool     `json:",string"`  //
	Score    float64
}

func main() {
	s := IT{"itcast", []string{"Go", "C++", "Java"}, false, 6.4};
	//buf, err := json.Marshal(s);
	buf, err := json.MarshalIndent(s, "", "    "); // indent最好4个空格，累容
	if (err != nil) {
		fmt.Println("err =", err);
		return;
	}
	fmt.Println("buf =", string(buf));
}

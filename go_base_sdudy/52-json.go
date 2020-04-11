package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	// 结构体变量一定要大写，否则获取不到
	Company  string
	Subjects []string
	IsOK     bool
	Score    float64
}

/**
结构体/map-=json.Marshal=>[]byte-=string()=>string
 */
func main() {
	s := IT{"itcast", []string{"Go", "C++", "Java"}, false, 6.4};
	//buf, err := json.Marshal(s);
	buf, err := json.MarshalIndent(s, "", "    "); // indent最好4个空格，累容
	// buf是byte类型
	if (err != nil) {
		fmt.Println("err =", err);
		return;
	}
	fmt.Println("buf =", string(buf));

	// 通过map生成json
	map1 := make(map[string]interface{}, 4);
	map1["Company"] = "q";
	map1["Subjects"] = 13;
	map1["IsOK"] = IT{"itcast", []string{"Go", "C++", "Java"}, false, 6.4};
	map1["a"] = [...]int{1, 3, 6, 9};
	buf, err = json.Marshal(map1);
	if (err != nil) {
		fmt.Println("err =", err);
		return;
	}
	fmt.Println("buf =", string(buf));
}

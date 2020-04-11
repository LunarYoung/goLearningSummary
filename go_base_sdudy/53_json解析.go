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
json-=[]byte()=->[]byte-=json.Unmarshal=->结构体/map
 */
func main() {
	s := "{\"Company\": \"itcast\",\"Subjects\": [\"Go\",\"C++\",\"Java\"],\"IsOK\": false,\"Score\": 6.4}";

	// 解析到结构体
	var tmp IT;
	fmt.Println("s =", s);
	err := json.Unmarshal([]byte(s), &tmp); // 只能传对象地址，指针还不行我去
	if (err != nil) {
		fmt.Println("err =", err);
		return;
	}
	fmt.Println(tmp);

	// 解析到map
	map1 := make(map[string]interface{}, 4);
	err = json.Unmarshal([]byte(s), &map1);
	if err != nil {
		fmt.Println("err =", err);
		return;
	}
	fmt.Println(map1);
	fmt.Println("str =", map1["Company"]); // 视频上说这里会报错，需要使用类型断言

	for key, value := range map1 {
		fmt.Printf("key = %v, value = %v...", key, value);
		switch value.(type) { // 类型断言的标准写法
		case string:
			fmt.Printf("the type of map[%s] is string", key);
			break;
		case int:
			fmt.Printf("the type of map[%s] is int", key);
			break;
		case []string:// 匹配不到
			fmt.Printf("the type of map[%s] is []string", key);
			break;
		case []interface{}:// 这才能匹配到
			fmt.Printf("the type of map[%s] is []interface", key);
			break;
		case bool:
			fmt.Printf("the type of map[%s] is bool", key);
			break;
		case float64:
			fmt.Printf("the type of map[%s] is float64", key);
			break;
		}
		fmt.Printf("\n");
	}
}

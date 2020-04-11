package main

import "fmt"

func main() {
	var map1 map[int]string;
	fmt.Println("map = ", map1);
	// map只有len，没有cap
	fmt.Printf("len = %d\n", len(map1));

	// 长度可选，相当于指定容量
	map2 := make(map[string]string, 10);
	map2["a"] = "A";
	fmt.Println("len(map2) = ", len(map2));
	map2["b"] = "B";
	fmt.Println("len(map2) = ", len(map2));
	map2["c"] = "C";
	fmt.Println("len(map2) = ", len(map2));
	fmt.Println("map2 = ", map2);
	fmt.Printf("%s", map2["d"]);
	map3 := map[int]string{1: "mike", 2: "Bob", 3: "Temple"};
	fmt.Println("map3 = ", map3);

	// 追加后，map底层自动扩容，和append类似，但是不会2倍扩

	for key, value := range map3 {
		fmt.Printf("%d = %s, ", key, value);
	}

	//delete(map3, 1);// 删除map3[1]
	value, ok := map3[1];
	if ok == true {
		fmt.Printf("%s", value);
	} else {
		fmt.Printf("该值不存在\n");
	}
}

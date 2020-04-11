package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args;
	if (len(list) != 2) {
		fmt.Println("usage: xxx file");
		return;
	}

	fileName := list[1];

	//
	info, infoErr := os.Stat(fileName);
	if (infoErr != nil) {
		fmt.Println("err =", infoErr);
		return;
	}

	//
	fmt.Println("name = ", info.Name());
	fmt.Println("size = ", info.Size());
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "http://www.bilibili.com/";
	resp, err := http.Get(url);
	if err != nil {
		fmt.Println("err =", err);
		return;
	}

	//
	defer resp.Body.Close();

	//
	fmt.Println("Status=", resp.Status);
	fmt.Println("Status=", resp.StatusCode);
	fmt.Println("Header=", resp.Header);
	fmt.Println("Body=", resp.Body);

	//
	buf := make([]byte, 1024*4);
	var tmp string;

	//
	for {
		n, readErr := resp.Body.Read(buf);
		if n == 0 {
			fmt.Println("readErr =", readErr);
			break;
		}

		tmp += string(buf[:n]);
	}

	//
	fmt.Println("tmp", tmp);
}

package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8010");
	if err != nil {
		fmt.Println("err =", err);
		return;
	}

	//
	defer conn.Close();

	//
	requestHeader := "GET /hello HTTP/1.1 Host: www.baidu.com User-Agent: Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.7.6)	Gecko/20050225 Firefox/1.0.1 Connection: Keep-Alive";

	// 先发请求包，服务器才会回响应包
	_, writeErr := conn.Write([]byte(requestHeader));
	if err != nil {
		fmt.Println("writeErr =", writeErr);
	}

	// 接收服务器回复的响应包
	buf := make([]byte, 1024*4);
	n, readErr := conn.Read(buf);
	if n == 0 {
		if readErr != nil {
			fmt.Println("readErr =", readErr);
		} else {
			fmt.Println("null response");
		}
		return;
	}

	//
	responseStr := string(buf[:n]);
	fmt.Printf("#%v#", responseStr);
}

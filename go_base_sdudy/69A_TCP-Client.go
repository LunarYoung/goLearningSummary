package main

import (
	"fmt"
	"net"
	"os"
)

const MAX_CHARS = 1024 // 最大字符串长度

func main() {
	// 主动连接服务器
	dialler, err := net.Dial("tcp", "127.0.0.1:8001");
	if (err != nil) {
		fmt.Println("err =", err);
		return;
	}

	// 调用完毕，关闭连接
	defer dialler.Close();

	//
	go receiveMsg(dialler);

	//
	for {
		//
		sendMsg(dialler);
	}

	////
	//var inputStr string;
	//for { // 无限次发送数据
	//	n, inputErr := fmt.Scan(&inputStr);
	//	if (inputErr != nil) {
	//		fmt.Println("err =", inputErr);
	//		return;
	//	}
	//	//fmt.Println(inputStr);
	//	if (n > Max_CHARS) {
	//		fmt.Println("OVERFLOW");
	//		return;
	//	}
	//
	//	//
	//	_, writeErr := dialler.Write([]byte(inputStr));
	//	if (writeErr != nil) {
	//		fmt.Println("err =", writeErr);
	//		return;
	//	}
	//	//fmt.Println("have sent");
	//}
}

func receiveMsg(conn net.Conn) {
	buf := make([]byte, MAX_CHARS);
	for {
		n, recErr := conn.Read(buf); // 接收服务器的请求
		if (recErr != nil) {
			fmt.Println("conn.Read err =", recErr);
			return;
		}
		fmt.Println(string(buf[:n])); // 接收到的内容转换为字符串，打印出来
	}
}

// 从键盘中读取内容，并发送到服务器
func sendMsg(conn net.Conn) {
	buf := make([]byte, MAX_CHARS);
	n, readErr := os.Stdin.Read(buf);
	if (readErr != nil) {
		fmt.Println("os.Stdin.Read err =", readErr);
		return;
	}

	// 把输入的内容发送到服务器
	conn.Write(buf[:n]);
}

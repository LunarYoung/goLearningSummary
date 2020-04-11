package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const MAX_CHARS = 1024;

func main() {
	// 提示输入文件
	fmt.Println("请输入需要传输的文件");
	var path string;
	fmt.Scan(&path);

	// 获取文件信息
	info, infoErr := os.Stat(path);
	if (infoErr != nil) {
		fmt.Println("os.Stat err =", infoErr);
		return;
	}

	// 连接服务器
	conn, connErr := net.Dial("tcp", "127.0.0.1:8002");
	if (connErr != nil) {
		fmt.Println("net.Dial err =", connErr);
		return;
	}
	defer conn.Close();

	// 给接收方，先发送文件名
	_, writeErr := conn.Write([]byte(info.Name()));
	if (writeErr != nil) {
		fmt.Println("conn.Write err =", writeErr);
		return;
	}

	// 接收对方回复，如果回复“ok”，说明对方准备好，可以发文件
	buf := make([]byte, MAX_CHARS);
	n, readErr := conn.Read(buf);
	if (readErr != nil) {
		fmt.Println("conn.Read err =", readErr);
	}

	//
	if (string(buf[:n]) == "ok") {
		// 发送文件内容
		sendFile(path, conn);
	}
}

func sendFile(path string, conn net.Conn) {

	// 以只读方式打开文件
	f, err := os.Open(path);
	if (err != nil) {
		fmt.Println("os.Open err =", err);
		return;
	}
	defer f.Close();
	buf := make([]byte, MAX_CHARS*MAX_CHARS);

	// 读文件内容
	for {
		n, readErr := f.Read(buf);
		if (readErr != nil) {
			if (readErr == io.EOF) {
				fmt.Println("文件发送完毕");
			} else {
				fmt.Println("f.Read err =", readErr);
			}
			return;
		}

		// 发送内容
		conn.Write(buf[:n]); // 给服务器发送内容
	}
}

package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const MAX_CHARS = 1024;

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8002");
	if (err != nil) {
		fmt.Println("listener err =", err);
		return;
	}
	defer listener.Close();

	// 阻塞等待用户连接
	conn, connErr := listener.Accept();
	if (connErr != nil) {
		fmt.Println("listener.Accept err =", connErr);
		return;
	}

	// 接收文件名称
	buf := make([]byte, MAX_CHARS);
	n, readErr := conn.Read(buf);
	if (readErr != nil) {
		fmt.Println("conn.Read =", err);
		return;
	}

	//
	if (n > 0) {
		fileName := string(buf[:n]);
		_, writeErr := conn.Write([]byte("ok"));
		if (writeErr != nil) {
			fmt.Println("conn.Write err =", writeErr);
			return;
		}

		// 接收文件内容
		receiveFile(conn, fileName);
	}
}

func receiveFile(conn net.Conn, fileName string) {
	buf := make([]byte, MAX_CHARS*MAX_CHARS);
	f, err := os.Create(fileName); // 这里因为是创建文件，所以不能用Open，而是用Create
	if (err != nil) {
		fmt.Println("os.Open err =", err);
		return;
	}
	defer f.Close();

	//
	for {
		n, readErr := conn.Read(buf);
		if (readErr != nil) {
			if (readErr == io.EOF) {
				fmt.Println("文件接收完毕");
			} else {
				fmt.Println("conn.Read err =", readErr);
			}
			return;
		}

		// 往文件写入内容
		_, fWriteErr := f.Write(buf[:n]);
		if (fWriteErr != nil) {
			fmt.Println("f.Write err =", fWriteErr);
			return;
		}
	}
}

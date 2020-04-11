package main

import (
	"fmt"
	"net"
	"strings"
)

const MAX_CHARS = 1024 // 最大字符串长度

func main() {
	// 监听
	listener, err := net.Listen("tcp", "127.0.0.1:8001");
	if (err != nil) {
		fmt.Println("err =", err);
		return;
	}

	defer listener.Close();
	for { // 无限次接收数据
		// 阻塞等待用户连接
		conn, connErr := listener.Accept();
		if (connErr != nil) {
			fmt.Println("err =", connErr);
			return;
		}

		//// 响应用户的请求
		//n, readErr := conn.Read(buf);
		//if (readErr != nil) {
		//	fmt.Println("err =", readErr);
		//	return;
		//}
		//
		//// 接收并显示字符串
		//recStr := string(buf[:n]);
		//if (recStr == "exit" || recStr == "quit") {
		//	conn.Close();
		//	break;
		//}
		//fmt.Println("buf =", recStr);
		go handle(conn);
	}
}

// 处理用户请求
func handle(conn net.Conn) {
	// 调用完毕，自动关闭conn
	defer conn.Close();

	// 获取客户端的网络地址信息
	addr := conn.RemoteAddr().String();
	fmt.Printf("{%s} connect successfully\n", addr);

	//
	buf := make([]byte, MAX_CHARS);
	for { // 无限次读取用户数据
		n, readErr := conn.Read(buf);
		if (readErr != nil) {
			fmt.Println("err =", readErr);
			return;
		}

		// 接收并显示字符串
		recStr := string(buf[:n-1]); // 最后的换行符也在其中，所以最后一个换行符不要
		if (recStr == "exit" || recStr == "quit") {
			fmt.Printf("{%s} exit\n", addr);
			return;
		}

		fmt.Printf("[%s]: %s\n", addr, recStr);

		// 把数据转换成大写，再给用户发送（被动发送消息）
		conn.Write([]byte(strings.ToUpper(recStr)));
	}

	//
}

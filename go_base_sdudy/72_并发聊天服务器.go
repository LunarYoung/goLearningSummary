package main

import (
	"fmt"
	"net"
	"time"
)

type Client struct {
	C    chan string // 用于发送数据的管道
	Name string      // 用户名
	Addr string      // 网络地址
}

// 保存在线用户  cliAttr=>Client
var onlineMap map[string]Client;

var message = make(chan string);

// 只要有消息来了，遍历map，给map每个成员都发送此消息
func Manager() {
	// 给map分配空间
	onlineMap = make(map[string]Client);
	for {
		msg := <-message; // 没有消息时，这里会阻塞

		// 遍历map，给map每个成员都发送此消息
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

// 向客户端发送消息
func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C { // 给当前客户端发送信息
		_, _ = conn.Write([]byte(msg));
	}
}

// 获取要发送的消息
func MakeMsg(cli Client, text string) (buf string) {
	return "【" + cli.Addr + "】" + cli.Name + "：" + text + "\n";
}

func HandleConn(conn net.Conn) { // 处理用户连接（用户上线了的处理）
	// 获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String();

	// 创建一个结构体
	cli := Client{make(chan string), cliAddr, cliAddr};

	// 把结构体添加到map
	onlineMap[cliAddr] = cli;

	// 新开一个协程，专门给客户端发送信息
	go WriteMsgToClient(cli, conn);

	// 广播某个人在线，所有客户端都能收到消息
	message <- MakeMsg(cli, "已登录");

	// 提示我是谁，这个只能自己收到
	cli.C <- MakeMsg(cli, "我在这里");

	//
	isQuit := make(chan bool);  // 对方是否主动退出
	hasData := make(chan bool); // 对方是否有数据发送

	// 接收用户发送过来的数据
	go func() {
		buf := make([]byte, 2048);
		for {
			n, readErr := conn.Read(buf);
			//if (readErr != nil) {
			//	fmt.Println("conn.Read error =", readErr);
			//}

			if (n == 0) { // 对方断开or出问题
				isQuit <- true;
				fmt.Println("conn.Read error =", readErr);
				return;
			}

			msg := string(buf[:n-1]); // nc 多一个换行

			if (len(msg) == 3 && msg == "who") { // 当收到“who”指令时，改为发送用户列表
				// 遍历map，给当前用户发送所有成员
				_, _ = conn.Write([]byte("User List:\n"));
				for _, tmp := range onlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\n";
					_, _ = conn.Write([]byte(msg));
				}
			} else if (len(msg) >= 8 && msg[:7] == "rename ") {
				// 更改用户名
				oldName := cli.Name;
				newName := msg[7:];
				cli.Name = newName;
				onlineMap[cliAddr] = cli;
				conn.Write([]byte("Client【" + cliAddr + "】的用户名更改" + oldName + "=>" + newName));
			} else {
				// 转发此内容
				fmt.Printf(MakeMsg(cli, msg));
				message <- MakeMsg(cli, msg);
			}

			hasData <- true; // 代表有数据
		}
	}()

	// 做个死循环，不要让方法结束
	for {
		// 通过select来检测channel的流动
		select {
		case <-isQuit:
			delete(onlineMap, cliAddr);     // 将当前用户从map中移除
			message <- MakeMsg(cli, "已退出"); // 广播下线
			return;
			//
		case <-hasData:
			break;
		case <-time.After(time.Second * 60):     // 60s后
			delete(onlineMap, cliAddr);          // 将当前用户从map中移除
			message <- MakeMsg(cli, "time out"); // 广播下线
			break;
		}

	}
}

func main() {
	listener, listenErr := net.Listen("tcp", ":8010");
	if (listenErr != nil) {
		fmt.Println("net.Listen error =", listenErr);
		return;
	}

	defer listener.Close();

	// 该协程用于转发消息，只要有消息来了，遍历map，给map每个成员都发送此消息
	go Manager();

	// 主协程，循环阻塞等待用户连接
	for {
		conn, connErr := listener.Accept();
		if (connErr != nil) {
			fmt.Println("listener.Accept =", connErr);
			continue;
		}

		go HandleConn(conn); // 处理用户连接
	}

}

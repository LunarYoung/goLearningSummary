package main

import (
	"fmt"
	"net"
)

const (
	MAX_SIZE = 1024
)

func main() {
	listener, err := net.Listen("tcp", ":8011");
	if err != nil {
		fmt.Println("err =", err);
		return;
	}

	defer listener.Close();

	//
	conn, connErr := listener.Accept();
	if connErr != nil {
		fmt.Println("err =", connErr);
		return;
	}
	defer conn.Close();

	//
	buf := make([]byte, MAX_SIZE);
	n, readErr := conn.Read(buf);
	if n == 0 {
		if readErr != nil {
			fmt.Println("readErr =", readErr);
			return;
		}
		fmt.Println("没有收到信息");
		return;
	}

	//fmt.Println(string(buf[:n]));
	fmt.Printf("#%v#", string(buf[:n]));
}

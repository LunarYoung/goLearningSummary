package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
设备文件：
屏幕（标准输出设备）fmt.Println() 往标准输出设备写内容
键盘（彼岸准输入设备）fmt.Scan() 从标准输入设备读取内容

磁盘文件，放在存储设备上的文件
1）文本文件		以记事本打开，能看到内容（不是乱码）
2）二进制文件		以记事本打开，能看到内容（是乱码）

为什么需要文件？
掉电丢失，程序结束，内存中的内容消失
文件放磁盘，程序结束，文件还是存在
 */

/**
Create
NewFile
Open
OpenFile
 */

/**
写文件
 */
func WriteFile(path string) {
	f, err := os.Create(path);
	if (err != nil) {
		fmt.Println("err =", err);
		return;
	}

	defer f.Close();

	// 开始写入
	var buf string;
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i = %d\n", i);
		n, err := f.WriteString(buf); // TMD追加式写入
		if (err != nil) {
			fmt.Println("err =", err);
			return;
		}
		fmt.Printf("n = %d\n", n);
	}
}

/**
读文件
 */
func ReadFile(path string) {
	// 打开文件
	f, err := os.Open(path);
	if (err != nil) {
		fmt.Println("err =", err);
		return;
	}

	defer f.Close();

	// 开始读取
	//var buf []byte;
	buf := make([]byte, 1024*2);
	n, err := f.Read(buf);
	if (err != nil && err != io.EOF) { // 文件出错，同时没到结尾
		fmt.Println("err =", err);
		return;
	}
	fmt.Println("size =", n);

	//fmt.Printf("n = %d\n", n);
	fmt.Println(string(buf[:n]));
}

/**
文件按行读取
 */
func ReadLine(path string) {
	// 打开文件
	f, err := os.Open(path);
	if (err != nil && err != io.EOF) {
		fmt.Println("err =", err);
	}

	// 关闭文件
	defer f.Close();

	// 新建一个缓冲区，把内容先放在缓冲区
	r := bufio.NewReader(f);
	line := 1;
	for {
		buf, err := r.ReadBytes('\n'); // 这种方法也会把\n读取进去
		if (err != nil) {
			if (err == io.EOF) { // 文件已经结束
				break;
			}
			fmt.Println("err =", err);
		}

		fmt.Printf("%d\t%s", line, string(buf));
		line++;
	}
}

func main() {
	fmt.Println("Are You OK?"); // 往标准输出设备（屏幕）写内容

	// 标准设备文件，默认已经打开，用户可以直接使用
	os.Stdout.WriteString("Are You OK?"); // 等价的
	//os.Stdout.Close();                    // 关闭后无法输出，读取文件时一定删掉

	path := "./iofiles/demo.txt";
	//WriteFile(path);
	ReadFile(path);
	ReadLine(path);
}

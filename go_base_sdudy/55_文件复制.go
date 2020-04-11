package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Command Line: go run this.go 1.txt 2.txt

	list := os.Args; // 获取命令行参数
	if (len(list) != 3) {
		fmt.Println("Usage: xxx srcFile dstFile");
		return;
	}

	srcFileName := list[1];
	dstFileName := list[2];
	if (srcFileName == dstFileName) {
		fmt.Println("源文件和目的文件名字相同");
		return;
	}

	// 只读方式打开源文件
	sf, err1 := os.Open(srcFileName);
	if (err1 != nil) {
		fmt.Println("err1 =", err1);
	}

	// 新建目的文件
	df, err2 := os.Create(dstFileName);
	if (err2 != nil) {
		fmt.Println("err2 =", err2);
	}

	// 操作完毕，需要关闭文件
	defer sf.Close();
	defer df.Close();

	// 核心处理，从源文件读取内容，往目的文件写，读多少写多少
	buf := make([]byte, 1024*4); // 4k大小的缓冲区
	for {
		n, err3 := sf.Read(buf);
		if (err3 != nil) {
			if (err3 == io.EOF) { // 文件读取完毕
				break;
			}
			fmt.Println("err3 =", err3);
		}
		m, err4 := df.Write(buf[:n]);
		if (err4 != nil) {
			fmt.Println("err4 =", err4);
		}
		fmt.Printf("m = %d", m);
	}
}

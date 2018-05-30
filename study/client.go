package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 通过tcp协议链接,本机1000端口
	conn, err := net.Dial("tcp", "127.0.0.1:1000")
	// 如果出现错误,说明链接失败
	if err != nil {
		fmt.Println("连接服务器端失败")
		fmt.Println(err.Error())
		os.Exit(0)
	}
	// 记得关闭 
	defer conn.Close()
	// 开始向服务器端发送 hello
	num, write_err := conn.Write([]byte("hello"))
	// 如果写入有问题 输出对应的错误信息
	if write_err != nil {
		fmt.Println(write_err.Error())
	}
	// 如果没有问题,显示对应的写入长度
	fmt.Println(num)
}



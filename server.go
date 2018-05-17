package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	//
	log.Println("Welcome go main!")
	// 开启服务
	StartServer()
}

// 处理错误
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// 开启服务
func StartServer() {
	// 使用tcp协议,监听本机1000端口
	netListen, err := net.Listen("tcp", "127.0.0.1:1000")
	//
	CheckError(err)
	//记得要关闭
	defer netListen.Close()
	//
	log.Println("Waiting for clients")
	// 循环接收
	for {
		// 等待链接,有链接过来 => 会赋值给 conn ,err
		conn, err := netListen.Accept()
		//
		CheckError(err)
		//
		log.Println(conn.RemoteAddr().String(), " tcp connect success")
		// 多个请求发送过来 => 并行处理
		go Handle(conn)
	}
}

// 处理
func Handle(conn net.Conn) {
	defer conn.Close()
	//创建一个缓冲*Reader 并读取对应的数据
	data, err := bufio.NewReader(conn).ReadString('\n')
	//如果数据读取完 err 会变成 EOF  这个并不是错误
	if err != nil && err != io.EOF {
		fmt.Println(err.Error())
	}
	log.Println(data)
}

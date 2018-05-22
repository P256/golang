package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	//
	log.Println("Welcome!")
	// 开启服务
	startServer()
}

// 处理错误
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// 开启服务
func startServer() {
	// 使用tcp协议,监听本机1000端口
	netListen, err := net.Listen("tcp", "0.0.0.0:1000")
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
		go handle3(conn)
	}
}

// 处理
func handle1(conn net.Conn) {
	defer conn.Close()
	//创建一个缓冲*Reader 并读取对应的数据
	data, err := bufio.NewReader(conn).ReadString('\n')
	//如果数据读取完 err 会变成 EOF  这个并不是错误
	if err != nil && err != io.EOF {
		fmt.Println(err.Error())
	}
	log.Println(data)
	//
}

//
func handle2(conn net.Conn) {
	//
	defer conn.Close()
	//
	buf := make([]byte, 1024)
	for {
		//
		reqLen, err := conn.Read(buf)
		fmt.Println(buf[0:reqLen])
		// 数据异常
		if err != nil && err != io.EOF {
			fmt.Println("err", err.Error())
		}
		// 遍历,转为16进制
		buffer := new(bytes.Buffer)
		for _, b := range buf[:reqLen] {
			s := strconv.FormatInt(int64(b&0xff), 16)
			if len(s) == 1 {
				buffer.WriteString("0")
			}
			buffer.WriteString(s)
		}
		// 转化为字符串
		fmt.Println(buffer.String())
		conn.Write([]byte(buffer.String()))
	}
}

//
func handle3(conn net.Conn) {
	//
	defer conn.Close()
	//
	data := make([]byte, 1024)
	// 读取数据-<ASCII码值>
	i, _ := conn.Read(data)
	// ASCII码方式读取
	dataStr1 := data[0:i]
	fmt.Println("ASCII码=>", dataStr1)
	// 字符串方式读取
	dataStr2 := string(data[0:i])
	fmt.Println("字符串=>", dataStr2)
	// 遍历,转为16进制
	buffer := new(bytes.Buffer)
	for _, b := range data[:i] {
		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(s)
	}
	// 转化为字符串
	data16 := buffer.String()
	fmt.Println("data16=>", data16)
	//conn.Write([]byte(data16))
	//
	data162 := []byte(data16)
	j, _ := conn.Read(data162)
	data162Str := string(data162[0:j])
	fmt.Println("data162ASCII码=>", data162[0:j])
	fmt.Println("data162=>", data162Str)
	//conn.Write([]byte("414243444546"))
}

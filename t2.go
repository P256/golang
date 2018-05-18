package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
)

var wg_server sync.WaitGroup
var clients = make(map[string]net.Conn)
var devices = make(map[string]net.Conn)

//var queues map[string]string
// 再使用make函数创建一个非nil的map，nil map不能赋值
var queues = make(map[string]string)

const (
	//绑定IP地址
	//ip = "127.0.0.1"
	ip = "192.168.1.79"
	//绑定端口号
	port = 110
)

func main() {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), port, ""})
	if err != nil {
		fmt.Println("监听端口失败:", err.Error())
		return
	}
	fmt.Println("已初始化连接，等待终端连接...")
	Server(listen)
}
func Server(listen *net.TCPListener) {
	//
	wg_server.Add(1)
	//接受Client消息
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("接受客户端连接异常:", err.Error())
			continue
		}
		//fmt.Println(strings.Contains("widuu", "wi")) //true
		//fmt.Println(strings.Contains("wi", "widuu")) //false

		var ipStr = conn.RemoteAddr().String()
		//fmt.Println(strings.Contains(ipStr, "102")) //true
		//strSplit := strings.Split(ipStr, ":")
		//fmt.Println(strSplit[1])
		ipStrSplit := strings.Split(ipStr, ":")
		port, err := strconv.Atoi(ipStrSplit[1])
		//
		if port < 500 {
			devices[conn.RemoteAddr().String()] = conn
			fmt.Println("设备端：" + ipStr + "上线")
			fmt.Println(devices)
		} else {
			clients[conn.RemoteAddr().String()] = conn
			fmt.Println("客户端：" + ipStr + "上线")
			fmt.Println(clients)
		}
		go Handle(conn)
	}
	wg_server.Done()
}
func Handle(conn net.Conn) {
	//conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	defer conn.Close()
	data := make([]byte, 1024)
	for {

/*//创建一个缓冲*Reader 并读取对应的数据
    data, err := bufio.NewReader(conn).ReadString('\n')
    //如果数据读取完 err 会变成 EOF  这个并不是错误
    if err != nil && err != io.EOF {
        fmt.Println(err.Error())
    }
    log.Println(data)*/
        
		i, err := conn.Read(data)
		fmt.Println("客户端", conn.RemoteAddr().String(), "发来数据:", string(data[0:i]))
		if err != nil {
			//fmt.Println("读取客户端数据错误:", err.Error())
			//fmt.Println("conn closed")
			DeleteClient(conn)
			break
		}
		if string(data[:i]) == "exit" {
			//conn.Write([]byte{'e', 'x', 'i', 't'})
			DeleteClient(conn)
			break
		}
		var no = string(data[0:1])
		var ip = ""
		if no == "1" {
			ip = "192.168.1.79:102"
		} else if no == "2" {
			ip = "192.168.1.79:101"
		} else {
			DeleteDevice(conn)
			break
		}
		//
		var share = devices[ip]
		fmt.Println(share)
		//
		if share != nil {
			fmt.Println("转发；", conn.RemoteAddr().String(), "=>", share.RemoteAddr().String(), "数据:", string(data[0:i]))
			share.Write(data[2:i])
			queues[no] = ip
			fmt.Println(queues)
		} else {
			fmt.Println("设备端：", conn.RemoteAddr().String(), "丢失")
			conn.Write([]byte("exit"))
		}
	}

}
func DeleteClient(conn net.Conn) {
	//fmt.Println(clients)
	fmt.Println("客户端" + conn.RemoteAddr().String() + " 已下线")
	delete(clients, conn.RemoteAddr().String())
	//fmt.Println("delete close conn")
	fmt.Println(clients)
	conn.Write([]byte("exit"))
	return
}
func DeleteDevice(conn net.Conn) {
	//fmt.Println(devices)
	fmt.Println("设备端" + conn.RemoteAddr().String() + " 已下线")
	delete(devices, conn.RemoteAddr().String())
	//fmt.Println("delete close conn")
	fmt.Println(devices)
	conn.Write([]byte("exit"))
	return
}

/*
https://studygolang.com/articles/2379
https://blog.csdn.net/Soooooooo8/article/details/70163475
https://blog.csdn.net/lanyang123456/article/details/78158255?locationNum=10&fps=1
https://blog.csdn.net/liwenjie0912/article/details/70187778
https://blog.csdn.net/wdy_2099/article/details/75073737

*/

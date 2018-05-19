/*
============================
手机	=> 服务 =>设备
============================
1.手机上线，获取设备在线列表
============================
设备序号	设备名称	设备状态
D50000	三星设备	可选/锁定/预约
============================
2.手机选择设备，分配设备=>设备配对
============================
手机序号	设备序号	设备状态
C10000	D50000	进行中/已完成
============================
3.手机操作技能，发动技能指令=>服务
============================
设备序号	设备技能
D50000	up
D50000	down
D50000	left
D50000	right
D50000	start
D50000	stop
D50000	rocker
============================
4.解析配对设备，发送技能指令=>设备
============================
*/
package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	//"time"
)

var waitGroup sync.WaitGroup
var clients = make(map[string]net.Conn)
var devices = make(map[string]net.Conn)
var queues = make(map[string]string)

//var clients1 [][]string

//var queues map[string]string

const (
	tcpIp = "127.0.0.1:1000"
)

func main() {
	// 绑定地址
	tcpAddr, err := net.ResolveTCPAddr("tcp", tcpIp)
	if err != nil {
		fmt.Println("绑定地址失败:", err.Error())
		return
	}
	// 监听端口
	listen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("监听端口失败:", err.Error())
		return
	}
	log.Println("已初始化连接,等待终端连接...")
	start(listen)
}
func start(listen *net.TCPListener) {
	//接受Client消息
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("接受客户端连接异常:", err.Error())
			continue
		}
		//fmt.Println(strings.Contains("widuu", "wi")) //true
		//fmt.Println(strings.Contains("wi", "widuu")) //false
		//
		var clientIp string = conn.RemoteAddr().String()
		ipStrSplit := strings.Split(clientIp, ":")
		port, err := strconv.Atoi(ipStrSplit[1])
		//
		if port < 500 {
			devices[clientIp] = conn
			log.Println("设备端:", clientIp+"上线")
			log.Println(devices)
		} else {
			clients[clientIp] = conn
			log.Println("客户端:", clientIp+"上线")
			log.Println(clients)
			conn.Write([]byte("deviceList"))
		}
		//
		waitGroup.Add(1)
		go handle(conn)
	}
	waitGroup.Wait()
}
func handle(conn net.Conn) {
	//conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	defer conn.Close()
	data := make([]byte, 1024)
	for {
		i, err := conn.Read(data)
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
		log.Println("客户端:", conn.RemoteAddr().String()+"发来数据:", string(data[0:i]))
		/*
			var deviceNo = string(data[0:1])
			if devices[deviceNo] {
				fmt.Println(devices[deviceNo])
			} else {
				DeleteDevice(conn)
			}
		*/
		var no = string(data[0:1])
		var ip = ""
		if no == "1" {
			ip = "127.0.0.1:102"
		} else if no == "2" {
			ip = "127.0.0.1:101"
		} else {
			DeleteDevice(conn)
			break
		}
		//
		var share = devices[ip]
		//fmt.Println(share)
		//
		if share != nil {
			log.Println("服务端:", conn.RemoteAddr().String()+"=>"+share.RemoteAddr().String()+"数据:"+string(data[0:i]))
			share.Write(data[2:i])
			queues[no] = ip
			fmt.Println(queues)
		} else {
			log.Println("设备端已下线")
			conn.Write([]byte("exit"))
		}
	}
	// 任务完成
	waitGroup.Done()
}
func DeleteClient(conn net.Conn) {
	//fmt.Println(clients)
	log.Println("客户端:", conn.RemoteAddr().String()+"下线")
	delete(clients, conn.RemoteAddr().String())
	log.Println(clients)
	conn.Write([]byte("exit"))
	return
}
func DeleteDevice(conn net.Conn) {
	//fmt.Println(devices)
	log.Println("设备端:", conn.RemoteAddr().String()+"下线")
	delete(devices, conn.RemoteAddr().String())
	log.Println(devices)
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

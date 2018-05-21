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
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
)

var waitGroup sync.WaitGroup
var clientConn = make(map[string]net.Conn)
var deviceConn = make(map[string]net.Conn)

var queues = make(map[string]string)

//var deviceList = make(map[int]string)
var deviceList []string

var groupList = make(map[string]interface{})

/*var deviceList = make(map[int]string)
 */

var numbers []int

//var clientConn1 [][]string

//var queues map[string]string

const (
	TCPIP = "192.168.1.79:1000"
)

func main() {
	// 绑定地址
	tcpAddr, err := net.ResolveTCPAddr("tcp", TCPIP)
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
		var clientIp = conn.RemoteAddr().String()
		ipStrSplit := strings.Split(clientIp, ":")
		port, _ := strconv.Atoi(ipStrSplit[1])
		//
		if port < 500 {
			deviceConn[clientIp] = conn
			//deviceList[port] = clientIp
			deviceList = append(deviceList, clientIp)
			log.Println("设备端:" + clientIp + "上线")
			log.Println(deviceConn)
		} else {
			clientConn[clientIp] = conn
			log.Println("客户端:" + clientIp + "上线")
			log.Println(clientConn)
			log.Println(deviceList)
			conn.Write([]byte("deviceList:"))
		}
		//
		waitGroup.Add(1)
		go handleTerminal(conn)
	}
	waitGroup.Wait()
}

func handleTerminal(conn net.Conn) {
	//conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	defer conn.Close()
	//
	data := make([]byte, 1024)
	for {
		//
		i, err := conn.Read(data)
		// 数据读取
		if err != nil && err != io.EOF {
			DeleteClient(conn)
			break
		}
		// 退出指令
		if string(data[:i]) == "exit" {
			DeleteClient(conn)
			break
		}
		// 打印输出
		/*jsonStr := string(data[0:i])
		log.Println("终端:"+conn.RemoteAddr().String()+"发来数据:", jsonStr)
		// 过滤标记
		jsonStr = strings.Replace(jsonStr, "'", "\"", -1)
		jsonStr = strings.Replace(jsonStr, "\n", "", -1)
		// var dat map[string]interface{}
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
			log.Println(dat)
		} else {
			log.Println(err)
		}
		// 设备序号
		//var deviceId = dat["deviceId"]
		//fmt.Println(deviceId, reflect.TypeOf(deviceId).String())
		// 输入参数1
		var deviceIp interface{}
		var deviceCmd string
		if _, ok := dat["deviceIp"]; ok {
			deviceIp = dat["deviceIp"]
		} else {
			DeleteClient(conn)
			break
		}
		// 输入参数2
		if _, ok2 := dat["action"]; ok2 {
			//deviceCmd = dat["action"]
			deviceCmd = "123"
		} else {
			DeleteClient(conn)
			break
		}*/

		// 2
		dataStr := string(data[0:i])
		log.Println("终端:"+conn.RemoteAddr().String()+"发来数据:", dataStr)
		// 过滤标记
		dataSplit := strings.Split(dataStr, ",")
		fmt.Println(dataSplit)
		fmt.Println(dataSplit[0])
		fmt.Println(dataSplit[1])

		// 输入参数1
		var deviceIp interface{}
		var deviceCmd string

		// 当前设备是否组队
		if _, ok := groupList[conn.RemoteAddr().String()]; ok {
			//fmt.Println("key存在")
			//
			var group = groupList[conn.RemoteAddr().String()]
			if group == deviceIp {
				//fmt.Println("val已存在")
				//
				var share = deviceConn["192.168.1.79:101"]
				//fmt.Println(share)
				//
				if share != nil {
					log.Println("服务端:", conn.RemoteAddr().String()+"=>"+share.RemoteAddr().String()+"数据:"+deviceCmd)
					share.Write([]byte(deviceCmd))
				} else {
					log.Println("设备端:", deviceIp, "已下线")
					conn.Write([]byte("exit"))
				}
			} else {
				log.Println("组队失败:", conn.RemoteAddr().String(), deviceIp, "提示: 设备已被其他锁定")
			}
		} else {
			// 设备是否可用

			//
			log.Println("组队成功:", conn.RemoteAddr().String(), deviceIp)
			// 组队
			groupList[conn.RemoteAddr().String()] = deviceIp
			log.Println(groupList)
		}

		/*var no = string(data[0:1])
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
		var share = deviceConn[ip]
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
		}*/
	}
	// 任务完成
	waitGroup.Done()
}
func DeleteClient(conn net.Conn) {
	//fmt.Println(clientConn)
	log.Println("客户端:" + conn.RemoteAddr().String() + "下线")
	delete(clientConn, conn.RemoteAddr().String())
	log.Println(clientConn)
	conn.Write([]byte("exit"))
	return
}
func DeleteDevice(conn net.Conn) {
	//fmt.Println(deviceConn)
	log.Println("设备端:" + conn.RemoteAddr().String() + "下线")
	delete(deviceConn, conn.RemoteAddr().String())
	log.Println(deviceConn)
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

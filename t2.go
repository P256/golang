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
	TCPIP = "0.0.0.0:1000"
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
	log.Println("服务端:已初始化连接,等待终端连接...")
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
		waitGroup.Add(1)
		go handleTerminal(conn)
	}
	waitGroup.Wait()
}

func handleTerminal(conn net.Conn) {
	//conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	defer conn.Close()
	//
	var remoteType = ""
	var remoteIp = conn.RemoteAddr().String()
	ipStrSplit := strings.Split(remoteIp, ":")
	port, _ := strconv.Atoi(ipStrSplit[1])
	// 据端口区分
	if port < 500 {
		remoteType = "device"
	} else {
		remoteType = "client"
	}
	// 终端类型
	if remoteType == "device" {
		deviceConn[remoteIp] = conn
		//deviceList[port] = terminalIp
		deviceList = append(deviceList, remoteIp)
		log.Println("设备端:" + remoteIp + "上线")
		log.Println("设备列:", deviceConn)
	} else {
		clientConn[remoteIp] = conn
		log.Println("客户端:" + remoteIp + "上线")
		log.Println("客户列:", clientConn)
		log.Println("设备列:", deviceList)
		var deviceStr = ""
		for i := 0; i < len(deviceList); i++ {
			// 以byte方式按字节遍历
			if deviceStr == "" {
				deviceStr = deviceList[i]
			} else {
				deviceStr = deviceStr + "," + deviceList[i]
			}
		}
		//log.Println(deviceStr)
		conn.Write([]byte("deviceList:" + deviceStr))
	}
	//
	data := make([]byte, 1024)
	for {
		//
		i, err := conn.Read(data)
		// 数据读取
		if err != nil && err != io.EOF {
			deleteConn(conn, remoteType, remoteIp)
			break
		}
		// 退出指令
		if string(data[:i]) == "exit" {
			deleteConn(conn, remoteType, remoteIp)
			break
		}
		// json方式
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

		// 字符串方式
		dataStr := string(data[0:i])
		log.Println("终端:" + remoteIp + "=>" + dataStr)
		// 过滤标记
		// TODO 下标问题
		dataSplit := strings.Split(dataStr, ",")
		//fmt.Println(reflect.TypeOf(dataSplit), dataSplit)
		// 验证输入参数
		var deviceIp = dataSplit[0]
		var deviceCmd = dataSplit[1]
		// 当前设备是否已组队
		groupIp := groupList[remoteIp]
		if groupIp == deviceIp {
			// 已组队
			log.Println("服务端:检测到当前与设备端已存在组队" + deviceIp)
			//
			var relay = deviceConn[deviceIp]
			//
			if relay != nil {
				log.Println("服务端:" + remoteIp + "=>" + relay.RemoteAddr().String() + "=>" + deviceCmd)
				relay.Write([]byte(deviceCmd))
			} else {
				log.Println("设备端:" + deviceIp + "已下线")
				conn.Write([]byte(deviceIp + ":exit"))
			}
		} else {
			// 检测未组队情况下所选择的设备是否可用
			var readOnly = false
			for _, group := range groupList {
				if deviceIp == group {
					readOnly = true
					break
				}
			}
			// 进行组队
			if !readOnly {
				log.Println("服务端:" + remoteIp + "<=>" + deviceIp + "组队成功")
				groupList[conn.RemoteAddr().String()] = deviceIp
			} else {
				log.Println("设备端:" + deviceIp + "已锁定")
				conn.Write([]byte(deviceIp + ":locked"))
				// TODO
			}
		}
		//
		log.Println(groupList)
	}
	// 任务完成
	waitGroup.Done()
}

// 删除连接
func deleteConn(conn net.Conn, remoteType string, remoteIp string) {
	if remoteType == "device" {
		log.Println("设备端:" + remoteIp + "下线")
		delete(deviceConn, remoteIp)
		log.Println(deviceConn)
	} else {
		log.Println("客户端:" + remoteIp + "下线")
		delete(clientConn, remoteIp)
		log.Println(clientConn)
	}
	conn.Write([]byte("exit"))
}

/*
https://studygolang.com/articles/2379
https://blog.csdn.net/Soooooooo8/article/details/70163475
https://blog.csdn.net/lanyang123456/article/details/78158255?locationNum=10&fps=1
https://blog.csdn.net/liwenjie0912/article/details/70187778
https://blog.csdn.net/wdy_2099/article/details/75073737

*/

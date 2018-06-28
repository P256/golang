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

//var deviceList []string
var deviceList = make(map[string]string)
var groupList = make(map[string]string)

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
	log.Println("服务端:", "已初始化连接,等待终端连接...")
	startListen(listen)
}
func startListen(listen *net.TCPListener) {
	//接受Client消息
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("接受客户端连接异常:", err.Error())
			continue
		}
		//
		waitGroup.Add(1)
		go handleTerminals(conn)
	}
	waitGroup.Wait()
}

func handleTerminals(conn net.Conn) {
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
		deviceList[ipStrSplit[1]] = remoteIp
		//deviceList = append(deviceList, remoteIp)
		log.Println("设备端:", remoteIp+"上线")
		log.Println("设备池:", deviceConn)
	} else {
		clientConn[remoteIp] = conn
		log.Println("客户端:", remoteIp+"上线")
		log.Println("客户池:", clientConn)
		//
		var deviceStr = ""
		//
		for key := range deviceList {
			// 以byte方式按字节遍历
			if deviceStr == "" {
				deviceStr = key
			} else {
				deviceStr = deviceStr + "," + string(key)
			}
		}
		//log.Println(deviceStr)
		conn.Write([]byte(deviceStr))
	}
	//
	log.Println("设备列:", deviceList)
	//
	data := make([]byte, 1024)
	for {
		// 读取数据
		i, err := conn.Read(data)
		// 数据异常
		if err != nil && err != io.EOF {
			deleteConn(conn, remoteType, remoteIp)
			break
		}
		// 退出指令
		if string(data[:i]) == "exit" {
			deleteConn(conn, remoteType, remoteIp)
			break
		}
		// 字符串方式
		dataStr := string(data[0:i])
		//
		if remoteType == "device" {
			log.Println("设备端:", remoteIp+"发来数据=>"+dataStr)
		} else {
			log.Println("客户端:", remoteIp+"发来数据=>"+dataStr)
		}
		//
		var deviceIp string
		// 是否已组队
		if _, ok := groupList[remoteIp]; ok {
			// 设备已组队
			deviceIp = groupList[remoteIp]
			log.Println("服务端:", "检测到当前客户端与设备端已组队"+deviceIp)
			//
			var relay = deviceConn[deviceIp]
			//
			if relay != nil {
				log.Println("服务端:", remoteIp+"=>"+relay.RemoteAddr().String()+"=>"+dataStr)
				relay.Write([]byte(dataStr))
			} else {
				log.Println("设备端:", deviceIp+"已下线")
				conn.Write([]byte("downLine"))
			}
		} else {
			//
			deviceNo := dataStr
			// 解析设备编号
			if v, ok := deviceList[deviceNo]; ok {
				deviceIp = v
			} else {
				// 设备列表找不到设备
				log.Println("服务端:", "设备端未找到")
				conn.Write([]byte("notFound"))
				break
			}
			// 设备是否可用
			var readOnly = false
			for _, value := range groupList {
				if deviceIp == value {
					readOnly = true
					break
				}
			}
			// 进行组队
			if !readOnly {
				log.Println("服务端:", remoteIp+"<=>"+deviceIp+"组队成功")
				groupList[conn.RemoteAddr().String()] = deviceIp
			} else {
				log.Println("设备端:", deviceIp+"已锁定")
				conn.Write([]byte("locked"))
			}
		}
		//
		log.Println("组队池:", groupList)
	}
	// 任务完成
	waitGroup.Done()
}

// 删除连接
func deleteConn(conn net.Conn, remoteType string, remoteIp string) {
	if remoteType == "device" {
		log.Println("设备端:", remoteIp+"下线")
		delete(deviceConn, remoteIp)
		log.Println("设备池:", deviceConn)
		// 删除deviceList
		for key, value := range deviceList {
			if remoteIp == value {
				delete(deviceList, key)
				break
			}
		}
		log.Println("设备列:", deviceList)
		// 删除groupList
		for key, value := range groupList {
			if remoteIp == value {
				delete(groupList, key)
				break
			}
		}
	} else {
		log.Println("客户端:", remoteIp+"下线")
		delete(clientConn, remoteIp)
		log.Println("客户池:", clientConn)
		delete(groupList, remoteIp)
	}
	//
	log.Println("组队池:", groupList)
	conn.Write([]byte("exit"))
}

/*
https://studygolang.com/articles/2379
https://blog.csdn.net/Soooooooo8/article/details/70163475
https://blog.csdn.net/lanyang123456/article/details/78158255?locationNum=10&fps=1
https://blog.csdn.net/liwenjie0912/article/details/70187778
https://blog.csdn.net/wdy_2099/article/details/75073737
*/

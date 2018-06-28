package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	// 字符串匹配
	fmt.Println(strings.Contains("widuu", "wi")) //true
	fmt.Println(strings.Contains("wi", "widuu")) //false

	// 字符串分割
	var remoteIp = "192.168.1.1:100"
	ipStrSplit := strings.Split(remoteIp, ":")
	fmt.Println(ipStrSplit[0], ipStrSplit[1], reflect.TypeOf(ipStrSplit[1]))

	// 字符串转int
	port, _ := strconv.Atoi(ipStrSplit[1])
	fmt.Println(port, reflect.TypeOf(port))

	// byte -> String（16进制）
	/*src := []byte("Hello")
	encodedStr := hex.EncodeToString(src)
	// 注意"Hello"与"encodedStr"不相等，encodedStr是用字符串来表示16进制
	fmt.Println(encodedStr)
	// String -> []byte
	test, _ := hex.DecodeString(encodedStr)
	fmt.Println(bytes.Compare(test, src)) // 0
	//
	test2, _ := hex.DecodeString(src)*/

	b := "Hello"
	fmt.Println(strconv.FormatInt(int64(b), 16))

}
